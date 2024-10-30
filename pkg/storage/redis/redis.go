package redis

import (
	"context"
	"fmt"
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

func ExampleClient() {
	rdb := NewStorage()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
