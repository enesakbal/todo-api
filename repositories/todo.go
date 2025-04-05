package repositories

import (
	"context"

	"github.com/enesakbal/todo-api/types"
)

type TodoRepository interface {
	CreateTodo(c context.Context, todo types.Todo, userID int) (types.Todo, error)
	GetTodosByUserID(c context.Context, userID int) ([]types.Todo, error)
}
