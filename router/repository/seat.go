package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type SeatRepository interface {
	SetRedis(key string, value interface{}, expiration time.Duration) (err error)
}

type seatRepository struct {
	Redis *redis.Client
}

func NewSeatRepository(redis *redis.Client) SeatRepository {
	return &seatRepository{
		Redis: redis,
	}
}

func (r *seatRepository) SetRedis(key string, value interface{}, expiration time.Duration) (err error) {
	return r.Redis.Set(context.Background(), key, value, expiration).Err()
}
