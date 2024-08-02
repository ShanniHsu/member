package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func Load() {
	vp := viper.New()
	vp.AddConfigPath(".")
	vp.SetConfigName("config.json")
	vp.SetConfigType("json")
	err := vp.ReadInConfig()
	if err != nil {
		log.Fatal("err: ", err)
		fmt.Println("讀取配置文件有錯誤!")
		return
	}
	name := vp.Get("name")
	fmt.Println(name)
}
