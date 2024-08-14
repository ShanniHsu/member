package migrate

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
	"member/pkg/storage"
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
	db := storage.InitStorage.GetDBConnect()
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202408131743",
			Migrate: func(tx *gorm.DB) error {
				// it's a good pratice to copy the struct inside the function,
				// so side effects are prevented if the original struct changes during the time
				type user struct {
					ID   int64 `gorm:"type:uuid;primaryKey;uniqueIndex"`
					Name string
				}
				return tx.Migrator().CreateTable(&user{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
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
	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("Migrate successfully!")
}
