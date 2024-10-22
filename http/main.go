package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a struct to hold the incoming data
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Function to handle requests for the root endpoint "/"
func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user User

		//Decode the JSON body into the user struct
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		//Respond with a message including the user's name and email
		responseMessage := fmt.Sprintf("Received user: Name: %s, Email: %s", user.Name, user.Email)
		fmt.Fprintf(w, responseMessage)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
func main() {

	// Map the "/" URL to the handleRequest function
	http.HandleFunc("/", handleRequest)

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
