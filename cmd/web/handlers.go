package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.mwhkdan.net/internal/models"
)

// Define a home handler function which writes a byte slice containing "Hello from Snippetbox" as the response body
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Set the content type to text/plain
	w.Header().Add("Server", "Go")
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.render(w, http.StatusOK, "home.tmpl", &templateData{Snippets: snippets})
}

// Define a snippetView handler that display a specific snippet
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Write the response body
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	app.render(w, http.StatusOK, "view.tmpl", &templateData{Snippet: &snippet})
}

// Define a snippetCreate handler that creates a new snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Set the content type to text/plain
	w.Header().Set("Content-Type", "text/plain")
	// Write the response body
	w.Write([]byte("Create a form for creting a new snippet..."))
}

func (app *application) snippetCreateTest(w http.ResponseWriter, r *http.Request) {
	// Set the content type to text/plain
	w.Header().Set("Content-Type", "text/plain")
	// Write the response body
	msg := fmt.Sprintf("Path: %s\n", r.PathValue("path"))
	w.Write([]byte(msg))
}

// Define a snippetCreatePost handler that save a new snippet
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Header().Add("Author", "dan")
	// Write the response body

	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji.\nBut slowly, slowly!\n\n-Kobayashi Issa"
	expires := 7
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
