package handler

import (
	"net/http"
	"strconv"

	"github.com/buskarion/todoapp-with-gin/service"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Status(ctx *gin.Context)
	GetTodos(ctx *gin.Context)
	GetTodosByID(ctx *gin.Context)
}

type handler struct {
	service service.Service
}

func NewHandler() Handler {
	return &handler{
		service: service.NewService(),
	}
}

func (h *handler) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "Server is up and running."})
}

func (h *handler) GetTodos(ctx *gin.Context) {
	var completed *bool
	completedParam := ctx.Query("completed")

	if completedParam != "" {
		parsedCompleted, err := strconv.ParseBool(completedParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query param."})
			return
		}
		completed = &parsedCompleted
	}

	todoList := h.service.GetAllTodos(completed)
	if len(todoList) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"msg": "The list is empty."})
		return
	}

	ctx.JSON(http.StatusOK, todoList)
}

func (h *handler) GetTodosByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	parsedID, err := strconv.Atoi(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	filteredTodo := h.service.FilterTodosByID(parsedID)

	if filteredTodo.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "ID not found."})
	}

	ctx.JSON(http.StatusOK, filteredTodo)
}
