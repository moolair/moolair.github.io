// Database connection setup
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL 드라이버 예제
)

func Connect(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return db, nil
}
