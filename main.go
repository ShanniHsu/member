package main

import (
	"member/config"
	"member/pkg/storage/migrate"
	"member/pkg/storage/mysql"
)

func main() {
	config.Init()
	mysql.Init()
	migrate.Init()
}
