package controller

import (
	"github.com/gin-gonic/gin"
	"member/router/service"
)

type AppController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserInfo(c *gin.Context)
	Logout(c *gin.Context)
}

type appController struct {
	userService service.User
}

func NewAppController(
	userService service.User) AppController {
	return appController{
		userService: userService,
	}
}
