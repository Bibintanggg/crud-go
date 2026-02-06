package database

import (
	"crud_go/internal/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUsername, dbPassword, dbName, dbPort,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	err = DB.AutoMigrate(
		&model.Book{},
	)

	if err != nil {
		panic("Can't connect to Model!")
	}

	fmt.Println("Succesfully connect to db")

	return nil

}
