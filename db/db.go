package db

import (
	"fmt"
	"log"
	"os"
	"todo-list/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)


func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, skipping...")
	}
	
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	fmt.Println("host:", host)
	fmt.Println("port:", port)
	fmt.Println("user:", user)
	fmt.Println("password:", password)
	fmt.Println("dbname:", dbname)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&models.Task{})

	return db

}
