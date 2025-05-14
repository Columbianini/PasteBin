package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	// Create a new ServeMux
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register the home handler function for the root URL path
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	mux.HandleFunc("GET /snippet/foo/{path}", app.snippetCreateTest)

	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}
