package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	// Set the content type to text/plain
	w.Header().Set("Content-Type", "text/plain")
	// Write the response body
	w.Write([]byte("Hello from Snippetbox"))
}

// Define a snippetView handler that display a specific snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Set the content type to text/plain
	w.Header().Set("Content-Type", "text/plain")
	// Write the response body
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// Define a snippetCreate handler that creates a new snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Set the content type to text/plain
	w.Header().Set("Content-Type", "text/plain")
	// Write the response body
	w.Write([]byte("Create a form for creting a new snippet..."))
}

func snippetCreateTest(w http.ResponseWriter, r *http.Request) {
	// Set the content type to text/plain
	w.Header().Set("Content-Type", "text/plain")
	// Write the response body
	msg := fmt.Sprintf("Path: %s\n", r.PathValue("path"))
	w.Write([]byte(msg))
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()
	// Register the home handler function for the root URL path
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/foo/{path}", snippetCreateTest)

	// Start the HTTP server on port 4000 and use the mux as the handler
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
