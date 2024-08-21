package migrate

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
	"member/pkg/storage"
	"member/sql"
)

func Init() {
	db := storage.InitStorage.GetDBConnect()
	migrateOption := gormigrate.DefaultOptions
	migrateOption.UseTransaction = true
	NewMigrate(db, migrateOption, sql.List)

	fmt.Println("Migrate successfully!")
}

func NewMigrate(db *gorm.DB, options *gormigrate.Options, list []*gormigrate.Migration) {
	m := gormigrate.New(db, options, list)
	m.InitSchema(func(tx *gorm.DB) (err error) {
		err = tx.AutoMigrate()
		if err != nil {
			return err
		}
		return nil
	})
	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	return
}
