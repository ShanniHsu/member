package service

import (
	"errors"
	"gorm.io/gorm"
	"member/models"
	"member/router/repository"
)

type Restaurant interface {
	GetRestaurants() (restaurants *[]models.Restaurant, err error)
}

type restaurantService struct {
	repo repository.Repo
}

func NewRestaurantService(repo repository.Repo) Restaurant {
	return restaurantService{
		repo: repo,
	}
}

func (s restaurantService) GetRestaurants() (restaurants *[]models.Restaurant, err error) {
	restaurants, err = s.repo.RestaurantRepository.GetRestaurants()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("Restaurants isn't found!")
		return
	}
	return
}
