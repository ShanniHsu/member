package models

import "time"

type Restaurant struct {
	ID        int64     `gorm:"id" json:"id"`
	Name      string    `gorm:"name" json:"name" comment:"餐廳名稱"`
	Type      int64     `gorm:"type" json:"type" comment:"餐廳品牌"`
	Address   string    `gorm:"address" json:"address" comment:"餐廳地址"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at" comment:"創建日期"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at" comment:"更新日期"`
}
