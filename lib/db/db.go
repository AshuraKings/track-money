package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func DbConn() (*sql.DB, error) {
	conn := os.Getenv("DATABASE_URL")
	if conn == "" {
		conn = "postgres://postgres:example@localhost:5432/track?sslmode=disable"
	}
	db, err := sql.Open("postgres", conn)
	return db, err
}
