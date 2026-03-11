package main

import (
	"log"
	"os"
	"forum/config"
	"forum/routes"
	"forum/handlers"
	"forum/repositories"
	"forum/services"
	"forum/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	db := config.ConnectDB()

	topicRepo := repositories.NewTopicRepository(db)
	postRepo := repositories.NewPostRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	userRepo := repositories.NewUserRepository(db)

	jwtService := services.NewJWTService(os.Getenv("JWT_SECRET"))
	userService := services.NewUserService(userRepo)
	topicService := services.NewTopicService(topicRepo, userRepo)
	postService := services.NewPostService(postRepo, topicRepo, userRepo)
	commentService := services.NewCommentService(commentRepo, postRepo, userRepo)

	userHandler := handlers.NewUserHandler(userService, jwtService)
	topicHandler := handlers.NewTopicHandler(topicService)
	postHandler := handlers.NewPostHandler(postService)
	commentHandler := handlers.NewCommentHandler(commentService)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authMiddleware := middleware.AuthMiddleware(jwtService)

	routes.RegisterRoutes(r, userHandler, topicHandler, postHandler, commentHandler, authMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server", err)
	}
}