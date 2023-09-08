package database

import (
	"log"

	"github.com/dassudip2001/webapp/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	log.Println("Connecting to database...")
	dsn := "admin:password@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}else{
		log.Println("Database connection successfully opened")
	}

	log.Println("Database connection successfully opened")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Run The Migrations")
	// add migrations here
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	Database = DbInstance{Db: db}
}
