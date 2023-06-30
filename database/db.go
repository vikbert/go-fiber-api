package database

import (
	"github.com/vikbert/go-fiber-api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDbInstance() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database!\n", err.Error())
		return
	}

	log.Println("connected to the database succeeded")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	errMigration := db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Order{},
	)

	if errMigration != nil {
		log.Fatal("DB migration failed!")
		return
	}

	Database = DbInstance{Db: db}
}
