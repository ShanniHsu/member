package main

import (
	"member/config"
	"member/pkg/shutdown"
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
	go router.Init()
	shutdown.Init()
}
