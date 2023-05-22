package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tykim96/webproject/src/config"
)

func InitDB() *sql.DB {
	cfg := config.LoadConfig()

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	log.Println("Connected to the database!")

	return db
}
