package main

import (
	"database/sql"
	"flag"
	"github.com/eunice99x/lets-go/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	//command line flags
	addr := flag.String("addr", ":4000", "HTTP listen address")
	dsn := flag.String("dsn", "root:root@/snippetbox?parseTime=true", "MYSQL DATA source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//dsn means data source name
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
	}

	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server", "addr", *addr)
	err = http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
