package router

import (
	"github.com/buskarion/todoapp-with-gin/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/status", handler.Status)

	return r
}
