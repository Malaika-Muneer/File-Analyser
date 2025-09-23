package DbConnection

import (
	"database/sql"
	"fmt"

	// main "go-mysql-app"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/malaika-muneer/File-Analyser/config"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	config.InitDB()
	// Check the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

}
