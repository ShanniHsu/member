package storage

import (
	"gorm.io/gorm"
	"member/pkg/storage/mysql"
)

type Storage struct {
	db *gorm.DB
}

var InitStorage Storage

func Init() {
	InitStorage.db = mysql.NewStorage()
}
