package models

import "time"

type User struct {
	ID        int64     `gorm:"id" json:"id"`
	Account   string    `gorm:"account" json:"account" comment:"帳號"`
	Password  string    `gorm:"password" json:"password" comment:"密碼"`
	Nickname  string    `gorm:"nickname" json:"nickname" comment:"綽號"`
	Status    int64     `gorm:"status" jason:"status" comment:"狀態"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}
