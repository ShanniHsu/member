package models

import "time"

type UserRestaurant struct {
	ID           int64     `gorm:"id" json:"id"`
	UserID       int64     `gorm:"user_id" json:"user_id" comment:"users.id"`
	RestaurantID int64     `gorm:"restaurant_id" json:"restaurant_id" comment:"restaurants.id"`
	CreatedAt    time.Time `gorm:"created_at" json:"created_at" comment:"創建日期"`
	UpdatedAt    time.Time `gorm:"updated_at" json:"updated_at" comment:"更新日期"`
}
