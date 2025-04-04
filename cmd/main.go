package main

import (
	middlewares "github.com/enesakbal/todo-api/api/middlewares"
	"github.com/enesakbal/todo-api/api/routes"
	"github.com/enesakbal/todo-api/bootstrap/application"
	"github.com/enesakbal/todo-api/bootstrap/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := application.App()

	env := app.Enviroment

	db := database.NewDatabase(env)
	defer app.CloseDBConnection()

	gin := gin.Default()

	gin.Use(middlewares.JsonLoggerMiddleware())

	routes.Setup(env, db, gin)

	gin.Run()
}
