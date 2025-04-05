package routes

import (
	"github.com/enesakbal/todo-api/api/controllers"
	"github.com/gin-gonic/gin"
)

func NewUserRoute(controllers controllers.UserController, group *gin.RouterGroup) {
	group.POST("/users", controllers.CreateUser)
	group.GET("/users", controllers.GetAllUsers)
	group.GET("/users/:user_id", controllers.GetUserByID)
}
