package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"weatherapp/internal/config"

	appHttp "weatherapp/internal/app/http"
)

type (
	App struct {
		config config.Config
		mux    *http.ServeMux
		server *http.Server
	}
)

func NewApp(ctx context.Context, config config.Config) (*App, error) {

	mux := http.NewServeMux()

	// Internal layer
	appHttp.RegisterInternalHandlers(mux)
	// END API ------------

	httpServerAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	// Merge components into app
	return &App{
		config: config,
		mux:    mux,
		server: &http.Server{Addr: httpServerAddr, Handler: mux},
	}, nil
}

func (a *App) Run(ctx context.Context, wg *sync.WaitGroup) error {
	// Start webserver
	log.Println("Starting HTTP-server")
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error starting server: %v", err)
		}
	}()

	log.Println("All components started")

	return nil
}

func (a *App) Close() error {
	err := a.server.Close()
	if err != nil {
		return err
	}
	return nil
}
