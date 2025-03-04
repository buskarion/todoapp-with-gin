package handler

import (
	"net/http"
	"strconv"

	"github.com/buskarion/todoapp-with-gin/service"
	"github.com/gin-gonic/gin"
)

func Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "Server is up and running."})
}

func GetTodos(ctx *gin.Context) {
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

	todoList := service.GetAllTodos(completed)
	if len(todoList) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"msg": "The list is empty."})
		return
	}

	ctx.JSON(http.StatusOK, todoList)
}

func GetTodosByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	parsedID, err := strconv.Atoi(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	filteredTodos := service.FilterTodosByID(parsedID)

	if len(filteredTodos) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "ID not found."})
	}

	ctx.JSON(http.StatusOK, filteredTodos)
}
