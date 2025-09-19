package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/malaika-muneer/File-Analyser/DbConnection"
	"github.com/malaika-muneer/File-Analyser/handlers"
)

func main() {

	DbConnection.ConnectDB()

	http.HandleFunc("/upload", handlers.UploadFile)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/signin", handlers.SignInHandler)

	// Create uploads directory if it doesn't exist
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err := os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create uploads directory: %v", err)
		}
	}

	fmt.Println("Server started at http://localhost:8005")
	log.Fatal(http.ListenAndServe(":8005", nil))

}
