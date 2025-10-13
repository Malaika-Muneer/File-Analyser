package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// var DB *sql.DB

func ConnectDB() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	return DB

}

// // GetDB returns the database instance
// func GetDB() *sql.DB {
// 	if DB == nil {
// 		log.Fatal("Database connection is not initialized. Call ConnectDB() first.")
// 	}
// 	return DB
// }
