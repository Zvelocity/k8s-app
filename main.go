package main

import (
	"fmt"
	"net/http"
)

// Handler function for the root URL
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Docker World!")
}

func main() {
	http.HandleFunc("/", helloHandler) // Set up a route to handle requests to the root URL
	// Start the web server on port 8080
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
