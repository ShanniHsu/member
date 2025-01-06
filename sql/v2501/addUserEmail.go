package v2501

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func AddUserEmail() (migrate *gormigrate.Migration) {
	return &gormigrate.Migration{
		ID: "202501061611",
		Migrate: func(tx *gorm.DB) (err error) {
			if err = tx.Exec("ALTER TABLE `users` ADD COLUMN `email` varchar(200) NOT NULL COMMENT 'E-mail' AFTER `nickname`").Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) (err error) {
			tx.Migrator().DropColumn("users", "email")
			return nil
		},
	}
}
