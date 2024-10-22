package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler function for the root endpoint "/"
// http.ResponseWriter writes a response to the client when they visit the page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to your first go HTTP server!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the About Page!")
}

func contactHomelander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the Contact Page")
}

func main() {
	// Map the "/" URL to the homeHandler function
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHomelander)

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
