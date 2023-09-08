package database

import (
	"log"
	"os"

	"github.com/dassudip2001/webapp/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database\n", err.Error())
		panic("Could not connect to the database")
		os.Exit(2)
	}

	log.Println("Database connection successfully opened")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Run The Migrations")
	// add migrations here
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	Database = DbInstance{Db: db}
}
