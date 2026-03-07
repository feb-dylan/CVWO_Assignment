package config

import (
	"fmt"
	"log"
	"os"
	"forum/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file" , err)
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable not set")
	}

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database" , err)
	}

	fmt.Println("Database connected")

	err = database.AutoMigrate(
			&models.User{},
		 	&models.Topic{},
		  	&models.Post{},
		   	&models.Comment{},
		)
		if err != nil {
			log.Fatal("Failed to migrate database" , err)
		}

		fmt.Println("Database migrated")
	return database
}
