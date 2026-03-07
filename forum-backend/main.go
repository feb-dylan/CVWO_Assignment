package main

import (
	"forum/config"
	"forum/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	r := gin.Default()

	routes.RegisterRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}
	log.Printf("Starting server on port %s", port)
	r.Run(":" + port)
}
