package middleware

import (
	"bytes"
	"compress/gzip"
	"github.com/gin-gonic/gin"
	"go-oo-demo/internal/pkg/logger"
	"io"
	"io/ioutil"
	"net/http"
)

// 默认http.Request.Body类型为io.ReadCloser类型,即只能读一次，读完后直接close掉,后续流程无法继续读取
func CopyRequestBody(c *gin.Context) {
	var maxMemory int64 = 64 << 20 // 64 MB

	var requestBody []byte
	isGzip := false
	safe := &io.LimitedReader{R: c.Request.Body, N: maxMemory}

	if c.GetHeader("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(safe)
		if err == nil {
			isGzip = true
			requestBody, _ = ioutil.ReadAll(reader)
		}
	}

	if !isGzip {
		requestBody, _ = ioutil.ReadAll(safe)
	}

	c.Request.Body.Close()
	bf := bytes.NewBuffer(requestBody)
	c.Request.Body = http.MaxBytesReader(c.Writer, ioutil.NopCloser(bf), maxMemory)
	c.Set(logger.RequestBodyKey, requestBody)

	c.Next()
}
