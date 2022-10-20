package main

import (
	"github.com/gin-gonic/gin"
	"go-grpc/errors_demo/api"
	_ "go-grpc/errors_demo/code"
)

func main() {
	e := gin.Default()
	e.GET("/user", api.GetUser)
	e.Run(":50001")
}
