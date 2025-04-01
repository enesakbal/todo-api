package services

import (
	"context"
	"database/sql"

	"github.com/enesakbal/todo-api/bootstrap/database"
	"github.com/enesakbal/todo-api/repositories"
	"github.com/enesakbal/todo-api/types"
)

type todoService struct {
	DBConnection sql.DB
}

func NewTodoService(db database.Database) repositories.TodoRepository {
	return &todoService{
		DBConnection: *db.DBConnection,
	}
}

func (t *todoService) CreateTodo(c context.Context, todo types.Todo) error {
	query := `
		INSERT INTO todos (
			user_id, 
			title, 
			description, 
			complete, 
			priority
		) VALUES (?, ?, ?, ?, ?)
	`

	_, err := t.DBConnection.Exec(query, todo.ID, todo.Title, todo.Description, todo.Complete, todo.Priority)

	if err != nil {
		return err
	}

	return nil
}
