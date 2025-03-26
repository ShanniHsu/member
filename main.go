package main

import (
	"member/pkg/storage/mgo"
)

func main() {
	//config.Init()
	//storage.Init()
	//migrate.Init()
	//jwt.InitJwt()
	//router.Init()
	mgo.Mongo()
}
