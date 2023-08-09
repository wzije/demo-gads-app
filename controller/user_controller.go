package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/wzcourses/demo-gads-app/model"
)

type UserController interface {
	Ping(c *gin.Context)
	Fetch(c *gin.Context)
}

type userController struct {
}

func (u userController) Ping(c *gin.Context) {
	c.String(200, "pong")
}

func (u userController) Fetch(c *gin.Context) {
	users := model.UserData
	c.JSON(200, users)
}

func NewUserController() UserController {
	return &userController{}
}
