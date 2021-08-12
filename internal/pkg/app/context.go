package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-oo-demo/internal/pkg/logger"
	"net/http"
)

type Context struct {
	*gin.Context
}

type Validator interface {
	Validate() error
}

func (c *Context) Query(key string) string {
	value, _ := c.GetQuery(key)
	return value
}

func (c *Context) Bind(data interface{}) error {
	if err := c.ShouldBind(data); err != nil {
		return err
	}

	if v, ok := data.(Validator); ok {
		return v.Validate()
	}

	return nil
}

func (c *Context) ToJson(data interface{}) {
	//c.JSON(http.StatusOK, data)

	buf, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	c.Set(logger.ResponseBodyKey, buf)
	c.Data(http.StatusOK, "application/json; charset=utf-8", buf)
	c.Abort()
}
