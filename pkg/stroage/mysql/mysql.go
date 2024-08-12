package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const (
	UserName     string = "root"
	Password     string = "password"
	NetWork      string = "tcp"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "test"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func Init() {

	//組合sql連線字串
	addr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, NetWork, Addr, Port, Database)
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
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)
}
