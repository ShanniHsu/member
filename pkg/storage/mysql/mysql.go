package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func NewStorage() (db *gorm.DB) {
	userName := viper.GetString("database.mysql.userName")
	password := viper.GetString("database.mysql.password")
	netWork := viper.GetString("database.mysql.netWork")
	host := viper.GetString("database.mysql.host")
	port := viper.GetInt("database.mysql.port")
	database := viper.GetString("database.mysql.database")
	maxLifetime := viper.GetInt("database.mysql.maxLifetime")
	maxOpenConns := viper.GetInt("database.mysql.maxOpenConns")
	maxIdleConns := viper.GetInt("database.mysql.maxIdleConns")

	//組合sql連線字串
	addr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True", userName, password, netWork, host, port, database)

	//連接MySQL
	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		log.Fatal("connection to mysql failed:", err)
	}

	//設定ConnMaxLifetime/MaxIdleConns/MaxOpenConns
	sqlDb, err := conn.DB()
	if err != nil {
		fmt.Println("get db failed:", err)
		log.Fatal("get db failed:", err)
	}
	sqlDb.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	sqlDb.SetMaxIdleConns(maxOpenConns)
	sqlDb.SetMaxOpenConns(maxIdleConns)
	fmt.Println("Connect DB Successfully!")
	return conn
}
