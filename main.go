package main

import (
	"member/pkg/webSocket"
	"os"
)

func main() {
	//config.Init()
	//storage.Init()
	//migrate.Init()
	//jwt.InitJwt()
	//router.Init()
	stuck := make(chan os.Signal)

	webSocket.ServerHTTP()

	<-stuck
}
