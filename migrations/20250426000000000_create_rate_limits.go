package migrations

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func MigrateRateLimits(ctx context.Context, conn *pgx.Conn) error {
	query := `
		CREATE TABLE IF NOT EXISTS rate_limits (
			is_paid BOOLEAN PRIMARY KEY,
			requests_per_minute INTEGER NOT NULL
		);
		INSERT INTO rate_limits (is_paid, requests_per_minute)
		VALUES 
			(FALSE, 10),
			(TRUE, 100)
		ON CONFLICT (is_paid) DO NOTHING;
	`
	_, err := conn.Exec(ctx, query)
	if err != nil {
		log.Printf("Error creating rate_limits table: %v", err)
		return err
	}
	log.Println("Rate_limits table created and seeded successfully")
	return nil
}

func RollbackRateLimits(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, "DROP TABLE IF EXISTS rate_limits;")
	if err != nil {
		log.Printf("Error dropping rate_limits table: %v", err)
		return err
	}
	log.Println("Rate_limits table dropped successfully")
	return nil
}
