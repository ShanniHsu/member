package main

import (
	"member/config"
	"member/pkg/jwt"
	"member/pkg/storage"
	"member/pkg/storage/migrate"
	"member/pkg/webSocket"
	"member/router"
)

func main() {
	go webSocket.Broadcast()
	config.Init()
	storage.Init()
	migrate.Init()
	jwt.InitJwt()
	router.Init()
}
