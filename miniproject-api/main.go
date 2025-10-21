package main

import (
	"github.com/gin-gonic/gin"
	_ "miniproject-api/docs" // Swagger docs
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

	"miniproject-api/internal/users"
	"miniproject-api/internal/threads"
	"miniproject-api/internal/posts"
	"miniproject-api/pkg/database"
)

// @title MiniProject API
// @version 1.0
// @description Twitter-like API with Users, Threads, Posts
// @host localhost:8080
// @BasePath /
func main() {
	// Init DB
	database.Init()
	users.AutoMigrate()
	threads.AutoMigrate()
	posts.AutoMigrate()

	r := gin.Default()

	// Users
	userService := users.NewService()
	userHandler := users.NewHandler(userService)
	userHandler.RegisterRoutes(r)

	// Threads
	threadService := threads.NewService()
	threadHandler := threads.NewHandler(threadService)
	threadHandler.RegisterRoutes(r)



	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
