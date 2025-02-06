package repository

import (
	"gorm.io/gorm"
	"member/models"
)

type UserRestaurantRepository interface {
	Create(userRestaurant *models.UserRestaurant) (err error)
}

type userRestaurantRepository struct {
	DB *gorm.DB
}

func NewUserRestaurantRepository(db *gorm.DB) UserRestaurantRepository {
	return userRestaurantRepository{
		DB: db,
	}
}

func (r userRestaurantRepository) Create(userRestaurant *models.UserRestaurant) (err error) {
	return r.DB.Create(userRestaurant).Error
}
