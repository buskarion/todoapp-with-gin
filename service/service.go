package service

import (
	"github.com/buskarion/todoapp-with-gin/db"
	"github.com/buskarion/todoapp-with-gin/entity"
)

type Service interface {
	GetAllTodos(completed *bool) []entity.Todo
	FilterTodosByID(id int) entity.Todo
	filterTodosByCompleted(completed *bool) []entity.Todo
}

type service struct {
	todos *[]entity.Todo
}

func NewService() Service {
	return &service{
		todos: db.BuildDB(),
	}
}

func (s *service) GetAllTodos(completed *bool) []entity.Todo {
	return s.filterTodosByCompleted(completed)
}

func (s *service) FilterTodosByID(id int) entity.Todo {
	for _, t := range *s.todos {
		if t.ID == id {
			return t
		}
	}
	return entity.Todo{}
}

func (s *service) filterTodosByCompleted(completed *bool) []entity.Todo {
	if completed == nil {
		return *s.todos
	}

	var filteredTodos []entity.Todo
	for _, t := range *s.todos {
		if t.Completed == *completed {
			filteredTodos = append(filteredTodos, t)
		}
	}

	return filteredTodos
}
