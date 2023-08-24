package main

import (
	"github.com/accmeboot/greenlight/internal/data"
	"github.com/accmeboot/greenlight/internal/jsonlog"
	"github.com/accmeboot/greenlight/internal/mailer"
	_ "github.com/lib/pq"
	"os"
	"sync"
)

const version = "1.0.0"

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

func main() {
	cfg := config{}
	cfg.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			logger.PrintFatal(err, nil)
		}
	}()

	logger.PrintInfo("database connection pool established", nil)

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	err = app.serve()
	logger.PrintFatal(err, nil)
}
