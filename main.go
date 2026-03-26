package main

import (
	"log"
	"os"
	"strings"

	"prog3bot/internal/handler"
	"prog3bot/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, relying on environment variables")
	}

	giphyAPIKey := os.Getenv("GIPHY_API_KEY")

	if giphyAPIKey == "" {
		log.Fatal("GIPHY_API_KEY must be set in environment variables.")
	}

	if strings.Contains(giphyAPIKey, "your_giphy_api_key") {
		log.Fatal("It looks like you haven't replaced the placeholder GIPHY_API_KEY in the .env file.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	giphyService := service.NewGiphyService(giphyAPIKey)
	r := handler.NewRouter(giphyService)
	app := r.SetupRoutes()

	log.Printf("Server started on :%s", port)
	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
