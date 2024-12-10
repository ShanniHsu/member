package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"member/models"
	"time"
)

// method
type UserRepository interface {
	GetUserByAccount(account string) (resp *models.User, err error)
	Create(user *models.User) (err error)
	SetRedis(key string, value string, expiration time.Duration) (err error)
	GetRedis(key string) (value string, err error)
	DeleteRedis(key string) (err error)
}

// the request that is applied the method
type userRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// use userRepository to return UserRepository(interface)
func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	return userRepository{
		DB:    db,
		Redis: redis,
	}
}

// Mothod of UserRepository(interface)
func (r userRepository) GetUserByAccount(account string) (resp *models.User, err error) {
	err = r.DB.Where("account = ?", account).First(&resp).Error
	return resp, err
}

func (r userRepository) Create(user *models.User) (err error) {
	return r.DB.Create(&user).Error
}

func (r userRepository) SetRedis(key string, value string, expiration time.Duration) (err error) {
	return r.Redis.Set(context.Background(), key, value, expiration).Err()
}

func (r userRepository) GetRedis(key string) (value string, err error) {
	return r.Redis.Get(context.Background(), key).Result()
}

func (r userRepository) DeleteRedis(key string) (err error) {
	return r.Redis.Del(context.Background(), key).Err()
}
