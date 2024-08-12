package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func Init() {
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
		return
	}
	//設定ConnMaxLifetime/MaxIdleConns/MaxOpenConns
	db, err1 := conn.DB()
	if err1 != nil {
		fmt.Println("get db failed:", err)
		return
	}
	db.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	db.SetMaxIdleConns(maxOpenConns)
	db.SetMaxOpenConns(maxIdleConns)
	fmt.Println("Connect DB Successfully!")
}
