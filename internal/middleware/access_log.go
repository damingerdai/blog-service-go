package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	 gin.ResponseWriter
	 body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: context.Writer}
		context.Writer = bodyWriter
		begin := time.Now().Unix()
		context.Next()
		end := time.Now().Unix()
		fmt.Printf("access log: request: %s, response: %s,method: %s, status_code: %d, begin time: %d, end time: %d", context.Request.PostForm.Encode(),bodyWriter.body.String(), context.Request.Method, bodyWriter.Status(), begin, end);
	}
}
