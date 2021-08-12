package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-oo-demo/interfaces/route"
	"go-oo-demo/interfaces/route/middleware"
	"go-oo-demo/internal/pkg/logger"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Server struct {
	config *ServerConfig
	handler http.Handler
	core *http.Server
}

func newServer(config *ServerConfig, db *gorm.DB, mode *Mode) *Server {
	server := new(Server)
	server.config = config

	gin.SetMode(mode.String())
	srv := gin.New()
	if mode.IsDebug() {
		srv.Use(gin.Logger())
	}
	srv.Use(middleware.TraceMiddleware)
	srv.Use(middleware.CopyRequestBody)
	srv.Use(middleware.Logger)

	router := route.New(srv, db)
	router.Init()

	server.handler = srv
	return server
}

func (server *Server) run(ctx context.Context) {

	server.core = &http.Server{
		Addr:         server.config.Port,
		Handler:      server.handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.WithContext(ctx).Printf("HTTP server is running at %s.", server.config.Port)
		err := server.core.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (server *Server) clear(ctx context.Context) {
	// http优雅关闭等待超时时长(单位秒)
	shutdownTimeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(shutdownTimeout))
	defer cancel()

	server.core.SetKeepAlivesEnabled(false)
	if err := server.core.Shutdown(ctx); err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}