package main

import (
	"fmt"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"time"
)

// 创建一个自定义中间件。

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")
		// 请求之前
		//Next应该只在中间件内部使用。它执行调用处理程序内部链中的挂起处理程序
		fmt.Println("中间件使用之前的处理1")
		c.Next()
		// 请求之后
		latency := time.Since(t)
		fmt.Println(latency, "2")
		// 访问我们发送的状态
		status := c.Writer.Status()
		fmt.Println(status, 3)
	}
}

func main() {
	e := gin.New()
	e.Use(Logger(), requestid.New())
	e.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		// it would print: "12345"
		fmt.Println(example)
		header := c.GetHeader("X-Request-ID")
		fmt.Println(header)
	})
	// Listen and serve on 0.0.0.0:8080
	e.Run(":50001")
}
