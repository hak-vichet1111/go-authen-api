package main

import (
	"go-authentication-api/handlers"
	"go-authentication-api/initializers"
	"go-authentication-api/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
) 

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDb()
    initializers.SyncDatabase()
}

func main() {
	app := gin.Default()
	// app.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Hello World!",
	// 	})
	// })

	// Public routes (do not require authentication)
	publicRoutes := app.Group("/public")
	{
		publicRoutes.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World!",
			})
		})

		// For our user authentication Login & Register
		publicRoutes.POST("/login", handlers.Login)
		publicRoutes.POST("/register", handlers.Register)
	}

	// Protected routes (require authentication)
	protectedRoutes := app.Group("/protected")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		protectedRoutes.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello Protected World!",
			})
		})
	}

	app.Run(":8080")
	log.Println("Server is running on port 8080")
}