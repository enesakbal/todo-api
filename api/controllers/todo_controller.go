package controllers

import (
	"net/http"

	"github.com/enesakbal/todo-api/types"
	"github.com/enesakbal/todo-api/usecases"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	Usecases usecases.TodoUsecases
}

func (tc *TodoController) CreateTodo(c *gin.Context) {
	var todo types.Todo

	err := c.ShouldBind(&todo)

	if err != nil {
		c.JSON(http.StatusBadRequest, "err")
		return
	}

	err = tc.Usecases.CreateTodo(c, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "errorr")
		return
	}

	c.JSON(http.StatusOK, "Todo created successfully")
}
