package mgo

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewStorage() (mdb *mongo.Client) {
	host := viper.GetString("database.mongo.host")
	port := viper.GetInt("database.mongo.port")
	url := fmt.Sprintf("mongodb://%s:%d", host, port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mdb, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	err = mdb.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB連線失敗: ", err)
	}
	fmt.Println("成功連線到 MongoDB")
	return
}
