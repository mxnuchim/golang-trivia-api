package database

import (
	"fmt"
	"log"
	"os"

	"github.com/mxnuchim/golang-trivia-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func ConnectDB() {

	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
		os.Exit(2)
	}

	log.Println("Database connected successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations...")
	db.AutoMigrate(&models.Fact{})

	DB = DBInstance{
		Db: db,
	}
}