package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // or postgres driver if you're using postgres
)

var DB *sql.DB

func InitDB() {
	// Read environment variables
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	fmt.Println("ENV values:",
		"USERNAME=", username,
		"PASSWORD=", password,
		"HOST=", host,
		"PORT=", port,
		"DBNAME=", dbname,
	)

	// Create DSN string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		username, password, host, port, dbname,
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	// check DB connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	fmt.Println("✅ Database connected successfully!")
}
