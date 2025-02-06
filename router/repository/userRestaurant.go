package repository

import "gorm.io/gorm"

type UserRestaurantRepository interface {
}

type userRestaurantRepository struct {
	DB *gorm.DB
}

func NewUserRestaurantRepository(db *gorm.DB) UserRestaurantRepository {
	return restaurantRepository{
		DB: db,
	}
}
