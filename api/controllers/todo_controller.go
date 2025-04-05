package controllers

import (
	"net/http"
	"strconv"

	"github.com/enesakbal/todo-api/types"
	"github.com/enesakbal/todo-api/usecases"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	Usecases usecases.TodoUsecases
}

func (tc *TodoController) CreateTodo(c *gin.Context) {

	userID, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid user ID")
		return
	}

	var todo types.Todo
	bindErr := c.ShouldBindJSON(&todo)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, bindErr.Error())
		return
	}

	todo, err = tc.Usecases.CreateTodo(c, todo, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error - "+err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (tc *TodoController) GetTodosByUserID(c *gin.Context) {

	//* /path?user_id=1234

	userID, err := strconv.Atoi(c.Request.URL.Query().Get("user_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid user ID - "+err.Error())
		return
	}

	todoList, err := tc.Usecases.GetTodosByUserID(c, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error - "+err.Error())
		return
	}

	c.JSON(http.StatusOK, todoList)
}
