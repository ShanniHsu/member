package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"member/models"
	create_user_restaurant "member/router/app/content/create-user-restaurant"
	delete_user_restaurant "member/router/app/content/delete-user-restaurant"
	get_user_restaurants "member/router/app/content/get-user-restaurants"
	"member/router/repository"
)

type UserRestaurant interface {
	GetPocketRestaurantList(ctx *gin.Context, req *get_user_restaurants.Request) (resp *get_user_restaurants.Response, err error)
	AddPocketRestaurant(ctx *gin.Context, req *create_user_restaurant.Request) (err error)
	DeletePocketRestaurant(ctx *gin.Context, req *delete_user_restaurant.Request) (err error)
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
		err = errors.New("Create Failed!")
		return
	}
	return
}

func (s userRestaurantService) DeletePocketRestaurant(ctx *gin.Context, req *delete_user_restaurant.Request) (err error) {
	var user = new(models.User)
	userCtx, exist := ctx.Get("user")
	if exist {
		user = userCtx.(*models.User)
	}
	req.UserID = user.ID

	checkList := new(get_user_restaurants.Request)
	checkList.ID = req.ID
	list, err := s.repo.UserRestaurantRepository.GetUserRestaurantFilter(checkList, req.UserID)
	if len(*list) == 0 {
		err = errors.New("ID isn't existed!")
		return
	}

	err = s.repo.UserRestaurantRepository.Delete(req)
	if err != nil {
		err = errors.New("Delete Failed!")
		return
	}
	return
}
