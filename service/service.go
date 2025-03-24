package service

import (
	"errors"
	"slices"

	"github.com/buskarion/todoapp-with-gin/db"
	"github.com/buskarion/todoapp-with-gin/entity"
)

type Service interface {
	GetAllTodos(completed *bool) []entity.Todo
	FilterTodosByID(id int) entity.Todo
	filterTodosByCompleted(completed *bool) []entity.Todo
	CreateTodo(todo entity.Todo) entity.Todo
	UpdateTodo(id int, updatedData entity.Todo) entity.Todo
	DeleteTodo(id int) error
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

func (s *service) CreateTodo(todo entity.Todo) entity.Todo {
	todo.ID = len(*s.todos) + 1
	todo.Completed = false
	*s.todos = append(*s.todos, todo)
	return todo
}

func (s *service) UpdateTodo(id int, updatedData entity.Todo) entity.Todo {
	todo := s.FilterTodosByID(id)
	todo.Task = updatedData.Task
	todo.Completed = updatedData.Completed
	return todo
}

func (s *service) DeleteTodo(id int) error {
	for i, t := range *s.todos {
		if t.ID == id {
			*s.todos = slices.Delete(*s.todos, i, i+1)
			return nil
		}
	}

	return errors.New("Todo not found")
}
