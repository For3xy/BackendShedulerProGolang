package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred while loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s  port=%s sslmode=disable", host, user, password, dbName, port)
	db, errSql := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errSql != nil {
		fmt.Println(err)
	}
	DB = db
	fmt.Println("Connection Successfully to the database")
}
