package migrate

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name    string
	Address string
}

type User struct {
	gorm.Model
	Name           string
	Age            int
	OrganizationID uint
}

func Init() {
	//test
	userName := viper.GetString("database.mysql.userName")
	password := viper.GetString("database.mysql.password")
	netWork := viper.GetString("database.mysql.netWork")
	host := viper.GetString("database.mysql.host")
	port := viper.GetInt("database.mysql.port")
	database := viper.GetString("database.mysql.database")

	//組合sql連線字串
	addr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True", userName, password, netWork, host, port, database)
	//連接MySQL
	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}

	m := gormigrate.New(conn, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{ID: "202408131743",
			Migrate:  nil,
			Rollback: nil,
		},
		// your migrations here
	})
	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&Organization{},
			&User{},
			// all other tables of you app
		)
		if err != nil {
			return err
		}

		if err := tx.Exec("ALTER TABLE users ADD CONSTRAINT fk_users_organizations FOREIGN KEY (organization_id) REFERENCES organizations (id)").Error; err != nil {
			return err
		}
		// all other constraints, indexes, etc...
		return nil
	})
	fmt.Println("Migrate successfully!")
}
