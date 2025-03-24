package router

import (
	"github.com/buskarion/todoapp-with-gin/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	h := handler.NewHandler()

	r.GET("/status", h.Status)
	r.GET("/todos", h.GetTodos)
	r.GET("/todos/:id", h.GetTodosByID)
	r.POST("/todos", h.CreateTodo)
	r.PUT("/todos/:id", h.UpdateTodo)
	r.PUT("/todos/:id", h.DeleteTodo)

	return r
}
