package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func NewStorage() (rdb *redis.Client) {
	host := viper.GetString("database.redis.host")
	port := viper.GetString("database.redis.port")
	password := viper.GetString("database.redis.password")
	dbName := viper.GetInt("database.redis.dbName")

	rdb = redis.NewClient(&redis.Options{
		Addr:     host + port,
		Password: password, // no password set
		DB:       dbName,   // use default DB
	})
	return rdb
}
