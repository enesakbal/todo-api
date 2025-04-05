package controllers

import (
	"net/http"
	"strconv"

	"github.com/enesakbal/todo-api/types"
	"github.com/enesakbal/todo-api/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Usecases usecases.UserUsecases
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user types.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, "err")
		return
	}

	err = uc.Usecases.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error - "+err.Error())
		return
	}

	c.JSON(http.StatusOK, "User created successfully")
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.Usecases.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error - "+err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := uc.Usecases.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error - "+err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
