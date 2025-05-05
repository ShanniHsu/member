package controller

import (
	"github.com/gin-gonic/gin"
	"member/router/service"
)

type AppController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserInfo(c *gin.Context)
	GetRestaurants(c *gin.Context)
	GetRestaurantList(c *gin.Context)
	GetPocketRestaurantList(c *gin.Context)
	AddPocketRestaurant(c *gin.Context)
	DeletePocketRestaurant(c *gin.Context)
	GetSeats(ctx *gin.Context)
	Logout(c *gin.Context)
}

type appController struct {
	userService       service.User
	restaurantService service.Restaurant
	userRestaurant    service.UserRestaurant
	seatService       service.SeatService
}

func NewAppController(
	userService service.User,
	restaurantService service.Restaurant,
	userRestaurantService service.UserRestaurant,
	seatService service.SeatService) AppController {
	return appController{
		userService:       userService,
		restaurantService: restaurantService,
		userRestaurant:    userRestaurantService,
		seatService:       seatService,
	}
}
