package main

import (
	"member/config"
	"member/pkg/newTicker"
	"member/pkg/storage"
	"member/pkg/storage/migrate"
	"member/router"
)

func main() {
	newTicker.Ticker()
	config.Init()
	storage.Init()
	migrate.Init()
	router.Init()
}
