package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	// Set the content type to text/plain
	w.Header().Add("Server", "Go")
	// Write the response body
	w.Write([]byte("Hello from Snippetbox"))
}

// Define a snippetView handler that display a specific snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Write the response body
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
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

// Define a snippetCreatePost handler that save a new snippet
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Header().Add("Author", "dan")
	w.WriteHeader(http.StatusCreated)
	// Write the response body
	w.Write([]byte("Save a new snippet..."))
}
