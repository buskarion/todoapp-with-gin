package service

import (
	"github.com/buskarion/todoapp-with-gin/db"
	"github.com/buskarion/todoapp-with-gin/entity"
)

var todos *[]entity.Todo = db.BuildDB()

func GetAllTodos() []entity.Todo {
	return *todos
}

func FilterTodosByID(id int) []entity.Todo {
	var filteredTodos []entity.Todo
	for _, t := range *todos {
		if t.ID == id {
			filteredTodos = append(filteredTodos, t)
		}
	}
	return filteredTodos
}
