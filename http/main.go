package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler function for the root endpoint "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to your first go HTTP server!")
}

func main() {
	// Map the "/" URL to the homeHandler function
	http.HandleFunc("/", homeHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
