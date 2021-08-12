package middleware

import (
	"github.com/gin-gonic/gin"
	"go-oo-demo/internal/pkg/logger"
	"mime"
	"net/http"
	"time"
)

func Logger(c *gin.Context) {
	path := c.Request.URL.Path
	method := c.Request.Method

	entry := logger.WithContext(logger.NewTagContext(c.Request.Context(), "__request__"))

	start := time.Now()
	fields := make(map[string]interface{})
	fields["method"] = method
	fields["url"] = c.Request.URL.String()
	fields["proto"] = c.Request.Proto
	fields["header"] = c.Request.Header
	fields["user_agent"] = c.GetHeader("User-Agent")
	fields["content_length"] = c.Request.ContentLength

	fields[logger.IPKey] = c.ClientIP()
	fields[logger.PathKey] = path

	if method == http.MethodPost || method == http.MethodPut {
		mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
		if mediaType != "multipart/form-data" {
			if v, ok := c.Get(logger.RequestBodyKey); ok {
				if b, ok := v.([]byte); ok {
					fields["body"] = string(b)
				}
			}
		}
	}
	c.Next()

	timeConsume := time.Since(start).Nanoseconds() / 1e6
	fields[logger.TimeConsumeKey] = timeConsume

	fields["res_status"] = c.Writer.Status()
	fields["res_length"] = c.Writer.Size()

	if v, ok := c.Get(logger.ResponseBodyKey); ok {
		if b, ok := v.([]byte); ok {
			fields["res_body"] = string(b)
		}
	}

	//fields[logger.UserIDKey] = ginx.GetUserID(c)

	entry.WithFields(fields).Infof("[http] %s-%s-%s-%d(%dms)",
		path, c.Request.Method, c.ClientIP(), c.Writer.Status(), timeConsume)
}
