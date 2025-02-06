package v2502

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateUserRestaurant() (migration *gormigrate.Migration) {
	return &gormigrate.Migration{
		ID: "202502061502",
		Migrate: func(tx *gorm.DB) (err error) {
			if err = tx.Exec("CREATE TABLE `user_restaurants` ( `id` int(11) unsigned NOT NULL AUTO_INCREMENT,\n  " +
				"`user_id` int(11) NOT NULL COMMENT 'users.id',\n  " +
				"`restaurant_id` int(11) NOT NULL COMMENT 'restaurants.id',\n  " +
				"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  " +
				"`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  " +
				"PRIMARY KEY (`id`),\n  " +
				"UNIQUE KEY `user_restaurant_unique` (`user_id`,`restaurant_id`)\n" +
				") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;").Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) (err error) {
			tx.Migrator().DropTable("user_restaurants")
			return nil
		},
	}
}
