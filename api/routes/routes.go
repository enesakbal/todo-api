package routes

import (
	"github.com/enesakbal/todo-api/api/controllers"
	"github.com/enesakbal/todo-api/bootstrap/database"
	enviroment "github.com/enesakbal/todo-api/bootstrap/env"
	"github.com/enesakbal/todo-api/services"
	"github.com/enesakbal/todo-api/usecases"
	"github.com/gin-gonic/gin"
)

func Setup(env *enviroment.Enviroment, db *database.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")

	todoRepository := services.NewTodoService(*db)
	todoUsecases := usecases.NewTodoUsecases(todoRepository)

	todoController := &controllers.TodoController{
		Usecases: todoUsecases,
	}

	NewTodoRoute(*todoController, publicRouter)
}
