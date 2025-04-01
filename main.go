package main

import (
	"member/channel"
	"member/config"
	"member/pkg/jwt"
	"member/pkg/storage"
	"member/pkg/storage/migrate"
	"member/router"
)

func main() {
	//channel.Channel()
	//channel.Ch()
	//channel.Foobar()
	//channel.MessageSend()
	channel.TestOne()
	channel.TestTwo()
	channel.TestThree()
	channel.TestOneAdvance()
	channel.TestTwoAdvance()
	channel.TestThreeAdvance()
	channel.Exam()
	config.Init()
	storage.Init()
	migrate.Init()
	jwt.InitJwt()
	router.Init()
}
