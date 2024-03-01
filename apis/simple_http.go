package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function for the "/" route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!") // Send "Hello, World!" as the response
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Handle GET requests
			fmt.Fprintf(w, "Hello, World! daas") // Send "Hello, World!" as the response
		} else if r.Method == http.MethodPost {
			// Handle POST requests
			fmt.Fprintf(w, "Received a POST request!") // Send acknowledgment for the POST request
		}
	})

	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
