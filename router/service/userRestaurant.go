package service

import (
	"github.com/gin-gonic/gin"
	"member/models"
	create_user_restaurant "member/router/app/content/create-user-restaurant"
	get_user_restaurants "member/router/app/content/get-user-restaurants"
	"member/router/repository"
)

type UserRestaurant interface {
	GetPocketRestaurantList(ctx *gin.Context, req *get_user_restaurants.Request) (resp *get_user_restaurants.Response, err error)
	AddPocketRestaurant(ctx *gin.Context, req *create_user_restaurant.Request) (err error)
}

type userRestaurantService struct {
	repo repository.Repo
}

func NewUserRestaurantService(repo repository.Repo) UserRestaurant {
	return userRestaurantService{
		repo: repo,
	}
}

func (s userRestaurantService) GetPocketRestaurantList(ctx *gin.Context, req *get_user_restaurants.Request) (resp *get_user_restaurants.Response, err error) {
	var user = new(models.User)

	userCtx, exist := ctx.Get("user")
	if exist {
		user = userCtx.(*models.User)
	}
	userID := user.ID
	resp, err = s.repo.UserRestaurantRepository.GetUserRestaurantFilter(req, userID)
	if err != nil {
		return
	}
	return
}

func (s userRestaurantService) AddPocketRestaurant(ctx *gin.Context, req *create_user_restaurant.Request) (err error) {
	var user = new(models.User)
	userCtx, exist := ctx.Get("user")
	if exist {
		user = userCtx.(*models.User)
	}

	userRestaurant := &models.UserRestaurant{
		UserID:       user.ID,
		RestaurantID: req.RestaurantID,
	}
	err = s.repo.UserRestaurantRepository.Create(userRestaurant)
	if err != nil {
		return
	}
	return
}
