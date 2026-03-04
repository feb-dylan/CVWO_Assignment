package main

import (
	"forum/config"
	"forum/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.RegisterRoutes(r, config.DB)

	port := ":8080"
	r.Run(port)
}
