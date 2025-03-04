package handler

import (
	"net/http"
	"strconv"

	"github.com/buskarion/todoapp-with-gin/db"
	"github.com/buskarion/todoapp-with-gin/entity"
	"github.com/gin-gonic/gin"
)

var todos *[]entity.Todo = db.BuildDB()

func Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "Server is up and running."})
}

func GetTodos(ctx *gin.Context) {
	caseDbIsEmpty(ctx)
	ctx.JSON(http.StatusOK, *todos)
}

func GetTodosByID(ctx *gin.Context) {
	caseDbIsEmpty(ctx)

	ID := ctx.Param("id")
	parsedID, err := strconv.Atoi(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	var filteredTodos []entity.Todo
	for _, t := range *todos {
		if t.ID == parsedID {
			filteredTodos = append(filteredTodos, t)
		}
	}

	if len(filteredTodos) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "ID not found."})
	}

	ctx.JSON(http.StatusOK, filteredTodos)
}

func caseDbIsEmpty(ctx *gin.Context) {
	if len(*todos) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"msg": "The list is empty."})
		return
	}
}
