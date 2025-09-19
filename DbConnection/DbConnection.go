package DbConnection

import (
	"database/sql"
	"fmt"

	// main "go-mysql-app"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	dsn := "root:NoPass@Ok032@tcp(localhost:3306)/fileanalyser"
	// Open a connection to the database
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

}
