package middleware

import (
	"bytes"
	"io/ioutil"
	"time"
	"yk/internal/app/robot"

	"github.com/gin-gonic/gin"
	z "go.uber.org/zap"
)

// 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		entry(c)
		defer exit(c)
		c.Next()
	}
}

// 请求进入日志
func entry(c *gin.Context) {
	// Recording request time
	c.Set("startExecTime", time.Now())

	// Read body data
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)

	// Write body back
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// Write request log
	robot.Logger().Info("entry", z.String("route", c.FullPath()), z.String("url", c.Request.URL.String()), z.String("body", string(bodyBytes)))
}

// 请求输出日志
func exit(c *gin.Context) {
	// Recording response time
	endExecTime := time.Now()

	// Get request time in gin.Context
	st, _ := c.Get("startExecTime")
	// Conversion interface to time.Time
	startExecTime := st.(time.Time)

	// Get response data (value, true) or (nil, false)
	data, _ := c.Get("resp_data_for_log")

	// Write response log
	robot.Logger().Info("exit", z.String("route", c.FullPath()), z.String("url", c.Request.URL.String()), z.Reflect("data", data), z.Float64("expend", endExecTime.Sub(startExecTime).Seconds()))
}
