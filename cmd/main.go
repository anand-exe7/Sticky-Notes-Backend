package main

import (
	"fmt"
	"net/http"
)

func main() {

	port := ":8080"
	
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"We gucci")
	})

	fmt.Println("Server Runninh on ",port)
	http.ListenAndServe(port,nil)
}