package service

import (
	"github.com/buskarion/todoapp-with-gin/db"
	"github.com/buskarion/todoapp-with-gin/entity"
)

var todos *[]entity.Todo = db.BuildDB()

func GetAllTodos(completed *bool) []entity.Todo {
	return filterTodosByCompleted(completed)
}

func FilterTodosByID(id int) entity.Todo {
	for _, t := range *todos {
		if t.ID == id {
			return t
		}
	}
	return entity.Todo{}
}

func filterTodosByCompleted(completed *bool) []entity.Todo {
	if completed == nil {
		return *todos
	}

	var filteredTodos []entity.Todo
	for _, t := range *todos {
		if t.Completed == *completed {
			filteredTodos = append(filteredTodos, t)
		}
	}

	return filteredTodos
}
