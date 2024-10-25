package main

import (
	"member/config"
	"member/pkg/argon2"
	"member/pkg/storage"
	"member/pkg/storage/migrate"
	"member/router"
)

func main() {
	argon2.TestPrint() //test the argon2
	config.Init()
	storage.Init()
	migrate.Init()
	router.Init()
}
