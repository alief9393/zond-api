// @title Zond Explorer API
// @version 1.0
// @description This is the API documentation for the Zond Explorer
// @host localhost:8080
// @BasePath /api

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"zond-api/internal/api"
	"zond-api/internal/config"
	"zond-api/internal/db"

	_ "zond-api/docs" // 👉 penting agar swag baca docs.go

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Failed to load .env file: %v", err)
	}

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to DB
	dbPool, err := db.NewDB(cfg.PostgresConn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbPool.Close()

	// JWT secret check
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatalf("JWT_SECRET is not set in .env or environment")
	}

	// Setup Gin router
	router := api.SetupRouter(dbPool, jwtSecret)

	// Register Swagger UI route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server in background
	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Graceful shutdown
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down server...")
}
