package repository

import (
	"gorm.io/gorm"
)

type RestaurantRepository interface {
}

type restaurantRepository struct {
	DB *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return restaurantRepository{
		DB: db,
	}
}
