package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func Connect(databaseURL string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test connection
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connected successfully")
	return &Database{Pool: pool}, nil
}

func (db *Database) Close() {
	db.Pool.Close()
	log.Println("Database connection closed")
}

func (db *Database) Health() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return db.Pool.Ping(ctx)
}

func (db *Database) GetPool() *pgxpool.Pool {
	return db.Pool
}

// RunMigrations executes database migrations for a specific service
func RunMigrations(db *Database, serviceName string) error {
	log.Printf("Running migrations for service: %s", serviceName)

	// TODO: Implement migration logic
	// This would typically use a migration tool like golang-migrate

	return nil
}

// Transaction helper
func (db *Database) WithTransaction(ctx context.Context, fn func(pgx.Tx) error) error {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	if err := fn(tx); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
