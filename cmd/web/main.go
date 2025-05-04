package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Pasrse the runtime configuration settings for the application
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Establish the dependencies for the handlers
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	app := &application{
		logger: logger,
	}

	// Start the HTTP server on port 4000 and use the mux as the handler
	logger.Info("Starting server", "addr", *addr)
	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
