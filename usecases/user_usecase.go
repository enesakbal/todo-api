package usecases

import (
	"context"

	"github.com/enesakbal/todo-api/repositories"
	"github.com/enesakbal/todo-api/types"
)

type UserUsecases struct {
	repository repositories.UserRepository
}

func NewUserUsecases(repository repositories.UserRepository) UserUsecases {
	usecase := UserUsecases{}

	usecase.repository = repository

	return usecase
}

func (usecase *UserUsecases) CreateUser(c context.Context, user types.User) error {
	return usecase.repository.CreateUser(c, user)
}

func (usecase *UserUsecases) GetAllUsers(c context.Context) ([]types.User, error) {
	return usecase.repository.GetAllUsers(c)
}
