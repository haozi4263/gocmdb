package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"grpc/gin_grpc/proto"
)

func main()  {
	 router := gin.Default()

	 router.GET("/proto", Proto)
	 router.Run(":8080")
}

func Proto(c *gin.Context)  {
	course := []string{"go","python"}
	teacher := proto.Teacher{Name: "jude", Course: course}
	c.ProtoBuf(http.StatusOK, &teacher)
}