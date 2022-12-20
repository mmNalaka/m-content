package database

import (
	"fmt"
	"time"

	"github.com/mmnalaka/m-content/config"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection(cfg *config.Postgres) (*sqlx.DB, error) {

	// Define database connection for PostgreSQL.
	db, err := sqlx.Connect("pgx", cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings:
	// 	- SetMaxOpenConns: the default is 0 (unlimited)
	// 	- SetMaxIdleConns: defaultMaxIdleConns = 2
	// 	- SetConnMaxLifetime: 0, connections are reused forever
	db.SetMaxOpenConns(cfg.MaxConnection)
	db.SetMaxIdleConns(cfg.MacIdle)
	db.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
