package main

import (
	"member/config"
	"member/pkg/jwt"
	"member/pkg/storage"
	"member/pkg/storage/migrate"
	"member/pkg/storage/redis"
	"member/router"
)

func main() {
	redis.ExampleClient()
	config.Init()
	storage.Init()
	migrate.Init()
	jwt.InitJwt()
	router.Init()
}
