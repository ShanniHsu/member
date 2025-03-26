package main

import (
	"member/config"
	"member/pkg/jwt"
	"member/pkg/storage"
	"member/pkg/storage/migrate"
	"member/router"
)

func main() {
	config.Init()
	storage.Init()
	migrate.Init()
	jwt.InitJwt()
	router.Init()
}
