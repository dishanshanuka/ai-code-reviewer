package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .env file එක load කරගැනීම (තියෙනවා නම්)
	if err := godotenv.Load(); err != nil {
		log.Println("Note: No .env file found, using default variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	// Health Check API Endpoint (Industrial Standard)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"message": "AI Code Reviewer Engine is running!",
		})
	})

	log.Printf("🚀 Server running on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}