package main

import (
	"member/config"
	"member/pkg/storage"
	"member/pkg/storage/migrate"
)

func main() {
	config.Init()
	storage.Init()
	migrate.Init()
}
