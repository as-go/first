package app

import (
	"net/http"

	"github.com/as-go/first/internal/handler"
	"github.com/as-go/first/internal/service"
	"github.com/as-go/first/internal/storage"
)

type app struct {
	address string
	handler http.Handler
}

type Config struct {
	Address          string
	ConnectionString string
}

func New(cfg Config) *app {
	store := storage.New(storage.Config{
		ConnectionString: cfg.ConnectionString,
	})

	userService := service.New(service.Config{
		Store: store,
	})

	h := handler.New(handler.Config{
		UserService: userService,
	})

	a := &app{
		address: cfg.Address,
		handler: h.Handler(),
	}

	return a
}

func (a *app) Start() {
	http.ListenAndServe(a.address, a.handler)
}
