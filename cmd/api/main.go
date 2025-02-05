package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/EmotionlessDev/todo-tasks/internal/data"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	logger *slog.Logger
	config config
	models data.Models
}

func main() {
	var cfg config

	// Parse the command-line flags
	flag.StringVar(&cfg.port, "port", ":4000", "HTTP network address")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "dsn", os.Getenv("TASKS_POSTGRES_DSN"), "Postgres connection DSN")

	flag.Parse()

	fmt.Println(cfg.db.dsn)
	// Init logger
	handler := slog.NewTextHandler(os.Stderr, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	errorLog := slog.NewLogLogger(handler, slog.LevelError)

	// Init the database connection pool
	db, err := openDB(cfg)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
	}
	defer db.Close()

	// Initialize a new instance of application containing the dependencies
	app := &application{
		logger: logger,
		config: cfg,
		models: data.NewModels(db),
	}

	// Add log and start the server with the servemux as the root handler
	logger.Info("Starting server on %s", slog.String("port", cfg.port))
	srv := &http.Server{
		Addr:     cfg.port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("Failed to start server", slog.Any("error", err))
	}

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
