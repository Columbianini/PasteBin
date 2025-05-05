package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"snippetbox.mwhkdan.net/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	// Pasrse the runtime configuration settings for the application
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Establish the dependencies for the handlers
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	defer db.Close()

	// Start the HTTP server on port 4000 and use the mux as the handler
	logger.Info("Starting server", "addr", *addr)
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	// Open a new connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Ping the database to check if it's reachable
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
