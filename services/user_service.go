package services

import (
	"context"
	"database/sql"

	"github.com/enesakbal/todo-api/bootstrap/database"
	"github.com/enesakbal/todo-api/repositories"
	"github.com/enesakbal/todo-api/types"
	"github.com/rs/zerolog/log"
)

type userService struct {
	DBConnection *sql.DB
}

func NewUserService(db database.Database) repositories.UserRepository {
	return &userService{
		DBConnection: db.DBConnection,
	}
}

func (t *userService) CreateUser(c context.Context, user types.User) error {
	query := `
		INSERT INTO users (
			user_id, 
			username, 
			email, 
			password_hash
		) VALUES (?, ?, ?, ?)
	`
	log.Info().Msg("Hello from Zerolog global logger")

	_, err := t.DBConnection.Exec(query, user.ID, user.Username, user.Email, user.Password)

	if err != nil {
		log.Error().Msg("Error creating user: " + err.Error())
		return err
	}

	return nil
}

func (t *userService) GetAllUsers(c context.Context) ([]types.User, error) {
	query := `
		SELECT * FROM users
	`

	rows, err := t.DBConnection.Query(query)

	if err != nil {
		log.Error().Msg("Error getting all users: " + err.Error())
		return nil, err
	}

	defer rows.Close()

	users := []types.User{}

	for rows.Next() {
		var user types.User

		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

		if err != nil {
			log.Error().Msg("Error scanning user: " + err.Error())
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
