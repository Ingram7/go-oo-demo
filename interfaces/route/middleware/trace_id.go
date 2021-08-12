package middleware

import (
	"github.com/gin-gonic/gin"
	"go-oo-demo/internal/pkg/logger"
	"go-oo-demo/internal/pkg/trace"
)

var t = new(trace.Trace)
// 跟踪ID中间件
func TraceMiddleware(c *gin.Context)  {


	// 优先从请求头中获取请求ID
	traceId := c.GetHeader("X-Request-Id")
	if traceId == "" {
		traceId = t.NewId()
	}

	ctx := t.Context(c.Request.Context(), traceId)
	ctx = logger.NewTraceIdContext(ctx, traceId)
	c.Request = c.Request.WithContext(ctx)
	c.Writer.Header().Set("X-Trace-Id", traceId)

	c.Next()
}
