package repositories

import (
	"context"

	"github.com/enesakbal/todo-api/types"
)

type TodoRepository interface {
	CreateTodo(c context.Context, todo types.Todo) error
}
