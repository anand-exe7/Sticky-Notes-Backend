package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		 log.Println("Error Getting the value from teh .env")
	}

	port := os.Getenv("PORT")
	
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"We gucci")
	})

	fmt.Println("Server Runninh on ",port)
	http.ListenAndServe(port,nil)
}