package services

import (
	"context"
	"database/sql"

	"github.com/enesakbal/todo-api/bootstrap/database"
	"github.com/enesakbal/todo-api/repositories"
	"github.com/enesakbal/todo-api/types"
	"github.com/rs/zerolog/log"
)

type todoService struct {
	DBConnection *sql.DB
}

func NewTodoService(db database.Database) repositories.TodoRepository {
	return &todoService{
		DBConnection: db.DBConnection,
	}
}

func (t *todoService) CreateTodo(c context.Context, todo types.Todo, userID int) (types.Todo, error) {
	query := `
		INSERT INTO todos (
			user_id, 
			title, 
			content, 
			complete, 
			priority
		) VALUES (?, ?, ?, ?, ?)
	`
	_, err := t.DBConnection.Exec(query, userID, todo.Title, todo.Content, todo.Complete, "medium")

	if err != nil {
		return types.Todo{}, err
	}

	todoX := &todo

	todoX.UserID = userID

	return *todoX, nil
}

func (t *todoService) GetTodosByUserID(c context.Context, userID int) ([]types.Todo, error) {
	query := `
		SELECT * FROM todos WHERE user_id = ?
	`

	rows, err := t.DBConnection.Query(query, userID)

	if err != nil {
		log.Error().Msg("Error creating todo: " + err.Error())
		return []types.Todo{}, err
	}

	defer rows.Close()

	todoList := []types.Todo{}

	for rows.Next() {
		var todo types.Todo

		err := rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.Complete, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt)

		if err != nil {
			return nil, err
		}

		todoList = append(todoList, todo)
	}

	return todoList, nil

}
