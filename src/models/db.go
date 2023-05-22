package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB establishes a connection to the MariaDB database.
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Set database connection settings (e.g., max connections, timeout)

	return db, nil
}
