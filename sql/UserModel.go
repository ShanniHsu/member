package sql

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"member/models"
)

func UserModel() (migrate *gormigrate.Migration) {

	return &gormigrate.Migration{
		ID: "1234",
		Migrate: func(tx *gorm.DB) (err error) {
			return tx.Migrator().CreateTable(&models.User{})
		},
		Rollback: func(tx *gorm.DB) (err error) {
			return tx.Migrator().DropTable("users")
		},
	}
}
