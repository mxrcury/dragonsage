package main

import (
	"fmt"
	"os"

	"github.com/mxrcury/dragonsage/internal/config"
	"github.com/mxrcury/dragonsage/internal/domain/http"
	v1 "github.com/mxrcury/dragonsage/internal/domain/http/v1"
	"github.com/mxrcury/dragonsage/internal/repository"
	"github.com/mxrcury/dragonsage/internal/service"
	"github.com/mxrcury/dragonsage/pkg/database"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	db, err := database.Init(cfg.DatabaseConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	repos := repository.NewRepositories(db)

	services := service.NewServices(repos)

	server := http.NewServer(cfg.ServerConfig)

	handlers := v1.NewHandlers(&v1.Deps{Router: server.Router, Services: services})
	handlers.Init()

	if err := server.Run(); err != nil {
		// add graceful shutdown
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
