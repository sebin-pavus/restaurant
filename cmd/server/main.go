package main

import (
	"log"
	"restaurant/internal/web"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	router := gin.Default()
	web.NewServer(router)

	router.Run("0.0.0.0:8080")
}
