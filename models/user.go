package models

import "time"

type User struct {
	ID        int64     `gorm:"id" json:"id"`
	Account   string    `gorm:"account" json:"account"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}
