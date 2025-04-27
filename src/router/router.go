package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/zGuiOs/poupeme-server/src/router/routes"
)

// Create and return an router with the routes
func Build() *chi.Mux {
	router := chi.NewRouter()

	return routes.Config(router)
}
