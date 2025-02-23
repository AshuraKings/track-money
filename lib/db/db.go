package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func DbConn() (*sql.DB, error) {
	conn := os.Getenv("DATABASE_URL")
	if conn == "" {
		conn = "postgres://neondb_owner:npg_cZesSba1D3TO@ep-lucky-sunset-a1b22guv-pooler.ap-southeast-1.aws.neon.tech/neondb?sslmode=require"
	}
	db, err := sql.Open("postgres", conn)
	return db, err
}
