package main

import (
	"member/config"
	"member/pkg/storage"
	"member/pkg/storage/migrate"
	"member/router"
)

func main() {
	config.Init()
	storage.Init()
	migrate.Init()
	//router.Init()

	//Test
	r := router.SetupRouter()
	r = router.PostUser(r)
	r.Run(":8080")
}
