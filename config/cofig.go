package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func Init() {
	vp := viper.New()
	//使用相對路徑
	vp.AddConfigPath(".") //所在目錄路徑
	vp.SetConfigName("config")
	err := vp.ReadInConfig()
	if err != nil {
		log.Fatal("err: ", err)
		fmt.Println("讀取配置文件有錯誤!")
		return
	}
	pwd, _ := os.Getwd() //目前所在位置
	fmt.Println(pwd)
	name := vp.GetString("name")
	fmt.Println(name)
}
