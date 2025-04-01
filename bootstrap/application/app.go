package application

import (
	"github.com/enesakbal/todo-api/bootstrap/database"
	enviroment "github.com/enesakbal/todo-api/bootstrap/env"
)

type Application struct {
	Database   *database.Database
	Enviroment *enviroment.Enviroment
}

func App() Application {
	app := &Application{}
	app.Enviroment = enviroment.NewEnviroment()
	app.Database = database.NewDatabase(app.Enviroment)

	return *app
}

func (app *Application) CloseDBConnection() error {
	return app.Database.CloseDBConnection()
}
