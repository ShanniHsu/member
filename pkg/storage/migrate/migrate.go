package migrate

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
	"member/pkg/storage"
	"member/sql"
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
	migrateOption := gormigrate.DefaultOptions
	migrateOption.UseTransaction = true
	NewMigrate(db, migrateOption, sql.List)
	//m := gormigrate.New(db, migrateOption, []*gormigrate.Migration{
	//	{
	//		ID: "202408131743",
	//		Migrate: func(tx *gorm.DB) error {
	//			// it's a good pratice to copy the struct inside the function,
	//			// so side effects are prevented if the original struct changes during the time
	//			return tx.Migrator().CreateTable(&models.User{})
	//		},
	//		//Migrate: func(tx *gorm.DB) (err error) {
	//		//	// it's a good pratice to copy the struct inside the function,
	//		//	// so side effects are prevented if the original struct changes during the time
	//		//	if err = tx.Exec("CREATE TABLE `users` (`id` int(11) unsigned NOT NULL AUTO_INCREMENT, \n" +
	//		//		"`account` varchar(200) NOT NULL COMMENT '帳號',\n" +
	//		//		"PRIMARY KEY (`id`)\n" +
	//		//		") ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;").Error; err != nil {
	//		//		fmt.Println("這有問題")
	//		//		//tx.Rollback()
	//		//		return err
	//		//	}
	//		//	return tx.Migrator().CreateTable(&models.User{})
	//		//},
	//		//Migrate: func(tx *gorm.DB) error {
	//		//	// it's a good pratice to copy the struct inside the function,
	//		//	// so side effects are prevented if the original struct changes during the time
	//		//	type user struct {
	//		//		ID   int64 `gorm:"type:uuid;primaryKey;uniqueIndex"`
	//		//		Name string
	//		//	}
	//		//	return tx.Migrator().CreateTable(&user{})
	//		//},
	//		Rollback: func(tx *gorm.DB) error {
	//			return tx.Migrator().DropTable("users")
	//		},
	//	},
	//	//{
	//	//	ID: "202408151634",
	//	//	Migrate: func(tx *gorm.DB) error {
	//	//		// it's a good pratice to copy the struct inside the function,
	//	//		// so side effects are prevented if the original struct changes during the time
	//	//		return tx.Migrator().CreateTable(&Organization{})
	//	//	},
	//	//},
	//	// your migrations here
	//})
	//
	//m.InitSchema(func(tx *gorm.DB) (err error) {
	//	//err = tx.AutoMigrate(
	//	//	//&Organization{},
	//	//	//&User{},
	//	//	&models.User{},
	//	//	// all other tables of you app
	//	//)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if err = tx.Exec("CREATE TABLE `users` (`id` int(11) unsigned NOT NULL AUTO_INCREMENT, \n" +
	//		"`account` varchar(200) NOT NULL COMMENT '帳號', \n" +
	//		"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, \n" +
	//		"`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, \n" +
	//		"PRIMARY KEY (`id`) \n" +
	//		") ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;").Error; err != nil {
	//		fmt.Println("這有問題")
	//		tx.Rollback()
	//		return err
	//	}
	//	//if err = tx.Exec("ALTER TABLE users ADD CONSTRAINT fk_users_organizations FOREIGN KEY (organization_id) REFERENCES organizations (id)").Error; err != nil {
	//	//	return err
	//	//}
	//	// all other constraints, indexes, etc...
	//	return nil
	//})

	//if err := m.Migrate(); err != nil {
	//	log.Fatalf("Migration failed: %v", err)
	//}
	fmt.Println("Migrate successfully!")
}

func NewMigrate(db *gorm.DB, options *gormigrate.Options, list []*gormigrate.Migration) {

	m := gormigrate.New(db, options, list)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	return
}
