package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define the port to listen on
	port := os.Getenv("PORT")
	if port == "" {
		port = "3008" // Default port if not specified
	}

	// Create a file server handler for serving static files
	fs := http.FileServer(http.Dir("static"))

	// Register the file server handler
	http.Handle("/", fs)

	// Start the web server
	fmt.Printf("Server starting on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
