package main

import (
	"github.com/gin-gonic/gin"
	"miniproject-api/internal/users"
)

func main() {
	r := gin.Default()

	userService := users.NewService()
	userHandler := users.NewHandler(userService)
	userHandler.RegisterRoutes(r)

	r.Run(":8080") // start server
}