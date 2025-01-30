package repository

import (
	"gorm.io/gorm"
	"member/models"
)

type RestaurantRepository interface {
	GetRestaurants() (restaurants *[]models.Restaurant, err error)
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
