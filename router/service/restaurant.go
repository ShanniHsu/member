package service

import "member/router/repository"

type Restaurant interface {
}

type restaurantService struct {
	repo repository.Repo
}

func NewRestaurantService(repo repository.Repo) Restaurant {
	return restaurantService{
		repo: repo,
	}
}
