package database

import (
	"github.com/vikbert/go-fiber-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDbInstance() {
	dsn := "host=localhost user=manfred password=manfred dbname=manfred port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the PostgreSQL!\n", err.Error())
		return
	}

	log.Println("connected to the PostgreSQL succeeded")
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
