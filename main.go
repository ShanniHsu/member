package main

import (
	"member/config"
	"member/pkg/stroage/mysql"
)

func main() {
	config.Init()
	mysql.Init()
}
