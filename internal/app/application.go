package app

import (
	"net/http"

	"github.com/ArjunDev17/course-content-service/internal/config"
	"github.com/ArjunDev17/course-content-service/internal/database"
)

type Application struct {
	Config *config.Config
	DB     *database.PostgreSQL
	Server *http.Server
}

func New() (*Application, error) {

	cfg := config.Load()

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
	}, nil
}
