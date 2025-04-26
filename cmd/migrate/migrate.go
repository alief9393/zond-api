package main

import (
	"context"
	"log"
	"os"

	"zond-api/migrations"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	dsn := os.Getenv("POSTGRES_CONN")
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer conn.Close(ctx)

	if err := migrations.MigrateUsers(ctx, conn); err != nil {
		log.Fatal("Users migration failed:", err)
	}
	if err := migrations.MigrateRateLimits(ctx, conn); err != nil {
		log.Fatal("Rate_limits migration failed:", err)
	}

	log.Println("All migrations completed successfully")
}
