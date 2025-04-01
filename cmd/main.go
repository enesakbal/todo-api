package main

import (
	"github.com/enesakbal/todo-api/api/routes"
	"github.com/enesakbal/todo-api/bootstrap/application"
	"github.com/enesakbal/todo-api/bootstrap/database"
	"github.com/gin-gonic/gin"
)

func main() {
	app := application.App()

	env := app.Enviroment

	db := database.NewDatabase(env)
	defer app.CloseDBConnection()

	gin := gin.Default()

	routes.Setup(env, db, gin)

	gin.Run(env.ServerAddress)
}
