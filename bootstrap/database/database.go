package database

import (
	"database/sql"
	"fmt"

	enviroment "github.com/enesakbal/todo-api/bootstrap/env"
)

type Database struct {
	Env          *enviroment.Enviroment
	DBConnection *sql.DB
}

func NewDatabase(env *enviroment.Enviroment) *Database {
	DBUser := env.DBUser
	DBPass := env.DBPass
	DBName := env.DBName
	DBHost := env.DBHost

	connectionSource := fmt.Sprintf("%v:%v@tcp(%v)/%v", DBUser, DBPass, DBHost, DBName)

	db, err := sql.Open("mysql", connectionSource)

	if err != nil {
		panic(err)
	}

	return &Database{
		env,
		db,
	}
}

func (database *Database) CloseDBConnection() error {

	return database.DBConnection.Close()

}
