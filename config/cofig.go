package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func Init() {
	SetDefault()
	//vp := viper.New()
	//使用相對路徑
	viper.AutomaticEnv()
	viper.AddConfigPath(".") //所在目錄路徑
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("err: ", err)
		fmt.Println("讀取配置文件有錯誤!")
		return
	}
	pwd, _ := os.Getwd() //目前所在位置
	fmt.Println(pwd)
	name := viper.GetString("name")
	username := viper.GetString("mysql.userName")
	fmt.Println(name)
	fmt.Println(username)

}

func SetDefault() {
	//mysql
	viper.SetDefault("mysql.userName", "Shanni")
	username := viper.GetString("mysql.userName")
	fmt.Println("SetDefault:", username)
}
