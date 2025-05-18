package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"snippetbox.mwhkdan.net/internal/models"

	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
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

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	app := &application{
		logger:        logger,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
		formDecoder:   formDecoder,
	}

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
