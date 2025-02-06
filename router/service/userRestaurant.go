package service

import "member/router/repository"

type UserRestaurant interface {
}

type userRestaurantService struct {
	Repo repository.Repo
}

func NewUserRestaurantService(repo repository.Repo) UserRestaurant {
	return restaurantService{
		repo: repo,
	}
}
