package storage

import (
	"gorm.io/gorm"
	"log"
	"member/pkg/storage/mysql"
)

type Storage struct {
	db *gorm.DB
}

var InitStorage Storage

func Init() {
	InitStorage.db = mysql.NewStorage()
}

// 表示都是使用指標(*Storage)
// 透過傳遞指標來操控同一個struct的實例
func (storage *Storage) GetDBConnect() *gorm.DB {
	if storage.db == nil {
		log.Fatal("Gat DB connect failed!")
	}
	return storage.db
}
