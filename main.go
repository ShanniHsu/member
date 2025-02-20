package main

import (
	"member/config"
	"member/pkg/deposit"
	"member/pkg/jwt"
	"member/pkg/storage"
	"member/pkg/storage/migrate"
	"member/router"
)

func main() {
	deposit.UseWorker()
	config.Init()
	storage.Init()
	migrate.Init()
	jwt.InitJwt()
	router.Init()
}
