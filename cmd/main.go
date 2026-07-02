package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sticky-notes-go-backend/internal/db"
	"sticky-notes-go-backend/internal/handler"
	"sticky-notes-go-backend/internal/routes"
	"sticky-notes-go-backend/internal/storer"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		 log.Println("Error Getting the value from teh .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	conn,err := db.NewDatabase(uri)
	if err != nil {
		log.Fatal("error connceting db")
	}
	
	storer := storer.NewDbStorer(conn.Client)
	
	handler := handler.NewHandler(storer)

	router := routes.NewRouter(handler)

	fmt.Println("Server Runninh on ",port)
	http.ListenAndServe(port,router)
}