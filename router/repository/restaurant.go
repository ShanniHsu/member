package repository

import (
	"gorm.io/gorm"
	"member/models"
	get_restaurants "member/router/app/content/get-restaurants"
)

type RestaurantRepository interface {
	GetRestaurants() (restaurants *[]models.Restaurant, err error)
	GetRestaurantList(parameter *get_restaurants.Request) (restaurants *[]models.Restaurant, err error)
}

type restaurantRepository struct {
	DB *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return restaurantRepository{
		DB: db,
	}
}

func (s restaurantRepository) GetRestaurants() (restaurants *[]models.Restaurant, err error) {
	err = s.DB.Model(&models.Restaurant{}).Find(&restaurants).Error
	return
}

func (s restaurantRepository) GetRestaurantList(parameter *get_restaurants.Request) (restaurants *[]models.Restaurant, err error) {
	query := s.DB.Model(&models.Restaurant{})
	if parameter.ID != 0 {
		query = query.Where("id = ?", parameter.ID)
	}
	if parameter.Type != 0 {
		query = query.Where("type = ?", parameter.Type)
	}
	err = query.Find(&restaurants).Error
	return
}
