package models

import "time"

type User struct {
	ID        int64     `gorm:"id" json:"id"`
	Account   string    `gorm:"account" json:"account" comment:"帳號"`
	Password  string    `gorm:"password" json:"password" comment:"密碼"`
	Nickname  string    `gorm:"nickname" json:"nickname" comment:"綽號"`
	Email     string    `gorm:"email" json:"email" comment:"E-mail"`
	Status    int64     `gorm:"status" json:"status" comment:"狀態"`
	Token     string    `gorm:"token" json:"token" comment:"Token"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at" comment:"創建日期"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at" comment:"更新日期"`
}
