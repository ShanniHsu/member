package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func Init() {
	SetDefault()
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
	fmt.Println("目前所在位置: ", pwd)

}

func SetDefault() {
	//mysql
	viper.SetDefault("database.mysql.userName", "root")
	viper.SetDefault("database.mysql.password", "password")
	viper.SetDefault("database.mysql.netWork", "tcp")
	viper.SetDefault("database.mysql.host", "127.0.0.1")
	viper.SetDefault("database.mysql.port", 3306)
	viper.SetDefault("database.mysql.database", "test")
	viper.SetDefault("database.mysql.maxLifetime", 10)
	viper.SetDefault("database.mysql.maxOpenConns", 10)
	viper.SetDefault("database.mysql.maxIdleConns", 10)
	//redis
	viper.SetDefault("database.redis.host", "localhost")
	viper.SetDefault("database.redis.port", "6379")
	viper.SetDefault("database.redis.password", "")
	viper.SetDefault("database.redis.dbName", 0)
}
