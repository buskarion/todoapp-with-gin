package router

import (
	"github.com/buskarion/todoapp-with-gin/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/status", handler.Status)
	r.GET("/todos", handler.GetTodos)
	r.GET("/todos/:id", handler.GetTodosByID)

	return r
}
