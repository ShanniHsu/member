package mgo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func Mongo() {
	url := "mongodb://localhost:27017"

	// 設定Context，並給予10秒的超時時間
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 連接Mongo
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	// 確認連線是否成功
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB連線失敗: ", err)
	}
	fmt.Println("成功連線到 MongoDB")

	defer client.Disconnect(ctx)
}
