package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Route represents a single route in the application.
type Route struct {
	URI      string
	Method   string
	Handler  func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}

// Config return all the routes to router
func Config(router *chi.Mux) *chi.Mux {
	routes := usersRoutes

	for _, route := range routes {
		router.MethodFunc(route.Method, route.URI, route.Handler)
	}

	return router
}
