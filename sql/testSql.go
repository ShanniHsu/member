package sql

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func TestSQL() (migrate *gormigrate.Migration) {
	return &gormigrate.Migration{
		ID: "09876",
		Migrate: func(tx *gorm.DB) (err error) {
			if err = tx.Exec("CREATE TABLE `tests` (`id` int(11) unsigned NOT NULL AUTO_INCREMENT, \n" +
				"`account` varchar(200) NOT NULL COMMENT '帳號', \n" +
				"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, \n" +
				"`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, \n" +
				"PRIMARY KEY (`id`) \n" +
				") ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;").Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) (err error) {
			tx.Migrator().DropTable("tests") //再確認
			return nil
		},
	}
}
