package main

import (
	"member/config"
	"member/pkg/jwt"
	"member/pkg/storage"
	"member/pkg/storage/migrate"
	"member/router"
	"member/router/service"
)

func main() {
	config.Init()
	storage.Init()
	migrate.Init()
	jwt.InitJwt()
	service.InitSeats()
	router.Init()
}
