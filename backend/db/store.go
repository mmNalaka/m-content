package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mmnalaka/m-content/db/sqlc"
)

// Connect to the the Postgres database
func OpenPostgresConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}

type SQLStore struct {
	db *sql.DB
	*sqlc.Queries
}

var Store *SQLStore

func CreteNewStore(db *sql.DB) {
	Store = &SQLStore{
		db:      db,
		Queries: sqlc.New(db),
	}
}
