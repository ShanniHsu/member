package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type SeatRepository interface {
	SetRedis(key string, value interface{}, expiration time.Duration) (err error)
	ScanRedis(cursor uint64, match string, count int64) (err error)
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

func (r *seatRepository) ScanRedis(cursor uint64, match string, count int64) (err error) {
	iter := r.Redis.Scan(context.Background(), cursor, match, count).Iterator()
	for iter.Next(context.Background()) {
		key := iter.Val()
		fmt.Println("key:", key)
	}
	if iter.Err() != nil {
		return
	}
	return
}
