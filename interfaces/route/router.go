package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-oo-demo/interfaces/controller"
	"go-oo-demo/internal/pkg/app"
	"gorm.io/gorm"
	"runtime"
)

type handler func(c *app.Context) (controller.Data, error)

type Router struct {
	engine *gin.Engine
	db     *gorm.DB
}

func New(engine *gin.Engine, db *gorm.DB) *Router {
	router := new(Router)
	router.engine = engine
	router.db = db

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	return router
}

func (router *Router) wrapper(handler handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		context := &app.Context{Context: c}

		defer func() {
			if r := recover(); r != nil {
				// 未知panic错误
				stack := deliverPanicStack(r)
				context.ToJson(controller.Response{Code: controller.StatusError, Message: controller.Message[controller.StatusError], Data: stack})
			}
		}()

		data, err := handler(context)
		code, message := parseError(err)
		context.ToJson(controller.Response{Code: code, Message: message, Data: data})
	}
}

func parseError(err error) (code int, message string) {
	code = controller.StatusOK
	message = controller.Message[controller.StatusOK]

	if err == nil {
		return
	}
	code = controller.StatusError
	message = err.Error()
	// 自定义错误
	error, ok := err.(controller.Error)
	if ok {
		code, message = error.Code, error.Message
		if message == "" {
			message = controller.Message[code]
		}
	}

	return
}

func deliverPanicStack(panic interface{}) string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	//if app.Config.Mode == gin.DebugMode {
	fmt.Sprintf("panic:\n%v\n %s\n", panic, string(buf[:n]))
	//}
	// 日志投递
	//middleware.Logger.WithFields(logrus.Fields{
	//	"panic": panic,
	//	"stack": string(buf[:n]),
	//}).Error("Panic")

	return fmt.Sprintf("panic:\n%v\n %s\n", panic, string(buf[:n]))
}
