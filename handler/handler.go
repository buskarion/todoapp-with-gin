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
	caseDbIsEmpty(ctx)
	ctx.JSON(http.StatusOK, service.GetAllTodos())
}

func GetTodosByID(ctx *gin.Context) {
	caseDbIsEmpty(ctx)

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

func caseDbIsEmpty(ctx *gin.Context) {
	todosList := service.GetAllTodos()
	if len(todosList) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"msg": "The list is empty."})
		return
	}
}
