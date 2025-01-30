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
	userService       service.User
	restaurantService service.Restaurant
}

func NewAppController(
	userService service.User,
	restaurantService service.Restaurant) AppController {
	return appController{
		userService:       userService,
		restaurantService: restaurantService,
	}
}
