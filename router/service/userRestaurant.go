package service

import (
	"errors"
	"member/models"
	create_user_restaurant "member/router/app/content/create-user-restaurant"
	delete_user_restaurant "member/router/app/content/delete-user-restaurant"
	get_restaurants "member/router/app/content/get-restaurants"
	get_user_restaurants "member/router/app/content/get-user-restaurants"
	"member/router/repository"
)

type UserRestaurant interface {
	GetPocketRestaurantList(user *models.User, req *get_user_restaurants.Request) (resp *get_user_restaurants.Response, err error)
	AddPocketRestaurant(user *models.User, req *create_user_restaurant.Request) (err error)
	DeletePocketRestaurant(user *models.User, req *delete_user_restaurant.Request) (err error)
}

type userRestaurantService struct {
	repo repository.Repo
}

func NewUserRestaurantService(repo repository.Repo) UserRestaurant {
	return userRestaurantService{
		repo: repo,
	}
}

func (s userRestaurantService) GetPocketRestaurantList(user *models.User, req *get_user_restaurants.Request) (resp *get_user_restaurants.Response, err error) {
	resp, err = s.repo.UserRestaurantRepository.GetUserRestaurantFilter(req, user.ID)
	if err != nil {
		err = errors.New("Get UserRestaurant Failed!")
		return
	}
	return
}

func (s userRestaurantService) AddPocketRestaurant(user *models.User, req *create_user_restaurant.Request) (err error) {
	parameter := &get_restaurants.Request{
		ID: req.RestaurantID,
	}
	_, err = s.repo.RestaurantRepository.GetRestaurantFilter(parameter)
	if err != nil {
		err = errors.New("Restaurant isn't existed!")
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

func (s userRestaurantService) DeletePocketRestaurant(user *models.User, req *delete_user_restaurant.Request) (err error) {
	checkList := &get_user_restaurants.Request{
		ID: req.ID,
	}

	res, err := s.repo.UserRestaurantRepository.GetUserRestaurantFilter(checkList, user.ID)
	if err != nil {
		err = errors.New("Get UserRestaurant Failed!")
		return
	}

	if res.TotalCount == 0 {
		err = errors.New("ID isn't existed!")
		return
	}

	err = s.repo.UserRestaurantRepository.DeleteByID(req.ID)
	if err != nil {
		err = errors.New("Delete Failed!")
		return
	}
	return
}
