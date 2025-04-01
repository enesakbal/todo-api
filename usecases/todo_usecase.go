package usecases

import (
	"context"

	"github.com/enesakbal/todo-api/repositories"
	"github.com/enesakbal/todo-api/types"
)

type TodoUsecases struct {
	repository repositories.TodoRepository
}

func NewTodoUsecases(repository repositories.TodoRepository) TodoUsecases {
	usecase := TodoUsecases{}

	usecase.repository = repository

	return usecase
}

func (usecase *TodoUsecases) CreateTodo(c context.Context, todo types.Todo) error {
	return usecase.repository.CreateTodo(c, todo)
}
