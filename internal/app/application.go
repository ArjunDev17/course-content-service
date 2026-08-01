package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ArjunDev17/course-content-service/internal/config"
	"github.com/ArjunDev17/course-content-service/internal/database"
)

type Application struct {
	Config *config.Config
	DB     *database.PostgreSQL
	Server *http.Server
	Logger *slog.Logger
}

func New() (*Application, error) {

	cfg := config.Load()

	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			nil,
		),
	)

	db, err := database.New(cfg.Database)
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr: ":" + cfg.App.Port,
	}

	return &Application{
		Config: cfg,
		DB:     db,
		Server: server,
		Logger: logger,
	}, nil
}

func (a *Application) Run() error {

	a.Logger.Info(
		"starting application",
		"service", a.Config.App.Name,
		"port", a.Config.App.Port,
	)

	go func() {

		if err := a.Server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {

			a.Logger.Error(
				"http server failed",
				"error", err,
			)

			os.Exit(1)
		}
	}()

	return a.waitForShutdown()
}

func (a *Application) waitForShutdown() error {

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	a.Logger.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {

		a.Logger.Error(
			"graceful shutdown failed",
			"error", err,
		)
	}

	a.DB.Close()

	a.Logger.Info("application stopped")

	return nil
}