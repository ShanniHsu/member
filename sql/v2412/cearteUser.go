package v2412

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateUser() (migrate *gormigrate.Migration) {
	return &gormigrate.Migration{
		ID: "202412191705",
		Migrate: func(tx *gorm.DB) (err error) {
			if err = tx.Exec("CREATE TABLE `users` (`id` int(11) unsigned NOT NULL AUTO_INCREMENT, \n" +
				"`account` varchar(200) NOT NULL COMMENT '帳號', \n" +
				"`password` varchar(200) NOT NULL COMMENT '密碼', \n" +
				"`nickname` varchar(200) NOT NULL COMMENT '綽號', \n" +
				"`status` int(11) NOT NULL COMMENT '狀態', \n" +
				"`token` varchar(500) NOT NULL COMMENT 'Token', \n" +
				"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, \n" +
				"`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, \n" +
				"PRIMARY KEY (`id`) \n" +
				") ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;").Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) (err error) {
			tx.Migrator().DropTable("users")
			return nil
		},
	}
}
