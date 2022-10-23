package main

import "github.com/gin-gonic/gin"

func main() {
	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery())                           // 所有请求都使用中间件
	e.Group("/v1").Use(gin.BasicAuth(gin.Accounts{"foo": "bar"})) // 组别全部使用中间件
	e.POST("/v2").Use()                                           // 特定的http请求使用中间件

}
