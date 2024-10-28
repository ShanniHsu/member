package storage

import (
	rdb "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"member/pkg/storage/mysql"
	"member/pkg/storage/redis"
)

type Storage struct {
	db  *gorm.DB
	rdb *rdb.Client
}

var InitStorage Storage

func Init() {
	InitStorage.db = mysql.NewStorage()
	InitStorage.rdb = redis.NewStorage()
}

// 表示都是使用指標(*Storage)
// 透過傳遞指標來操控同一個struct的實例
func (storage *Storage) GetDBConnect() *gorm.DB {
	if storage.db == nil {
		log.Fatal("Get DB connect failed!")
	}
	return storage.db
}

func (storage *Storage) GetRDBConnect() *rdb.Client {
	if storage.rdb == nil {
		log.Fatal("Get RDB connect failed!")
	}
	return storage.rdb
}
