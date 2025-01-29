package v2501

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateRestaurant() (migrate *gormigrate.Migration) {
	return &gormigrate.Migration{
		ID: "202501300254",
		Migrate: func(tx *gorm.DB) (err error) {
			if err = tx.Exec("CREATE TABLE `restaurants` (`id` int(11) unsigned NOT NULL AUTO_INCREMENT, \n" +
				"`name` varchar(200) NOT NULL DEFAULT '' COMMENT '餐廳名稱',\n  " +
				"`type` int(11) NOT NULL COMMENT '品牌類型',\n  " +
				"`address` varchar(200) NOT NULL DEFAULT '' COMMENT '餐廳地址',\n  " +
				"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  " +
				"`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  " +
				"PRIMARY KEY (`id`)\n" +
				") ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;").Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) (err error) {
			tx.Migrator().DropTable("restaurants")
			return nil
		},
	}
}
