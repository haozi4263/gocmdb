package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/server/egin"
)

func HttpServer() *egin.Component {
	Gin := egin.Load("server.http").Build()
	Gin.GET("/hello", func(c *gin.Context) {
		c.JSON(200, "Hello grpc_http for http")
	})
	return Gin
}
