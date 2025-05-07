package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type SeatRepository interface {
	SetRedis(key string, value interface{}, expiration time.Duration) (err error)
	ScanRedis(cursor uint64, match string, count int64) (value map[string]string, err error)
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

func (r *seatRepository) ScanRedis(cursor uint64, match string, count int64) (value map[string]string, err error) {
	iter := r.Redis.Scan(context.Background(), cursor, match, count).Iterator()
	value = make(map[string]string)
	for iter.Next(context.Background()) {
		key := iter.Val()
		status := ""
		status, err = r.Redis.Get(context.Background(), key).Result()
		if err != nil {
			return
		}
		value[key] = status
	}
	if iter.Err() != nil {
		return
	}
	return
}
