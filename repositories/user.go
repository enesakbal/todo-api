package repositories

import (
	"context"

	"github.com/enesakbal/todo-api/types"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user types.User) error
	GetAllUsers(ctx context.Context) ([]types.User, error)
}
