package migrations

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func MigrateUsers(ctx context.Context, conn *pgx.Conn) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			is_paid BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := conn.Exec(ctx, query)
	if err != nil {
		log.Printf("Error creating users table: %v", err)
		return err
	}
	log.Println("Users table created successfully")
	return nil
}

func RollbackUsers(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, "DROP TABLE IF EXISTS users;")
	if err != nil {
		log.Printf("Error dropping users table: %v", err)
		return err
	}
	log.Println("Users table dropped successfully")
	return nil
}
