package routes

import (
	"github.com/enesakbal/todo-api/api/controllers"
	"github.com/gin-gonic/gin"
)

func NewTodoRoute(controllers controllers.TodoController, group *gin.RouterGroup) {

	group.POST("/todo", controllers.CreateTodo)
}
