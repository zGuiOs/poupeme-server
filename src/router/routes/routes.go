package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zGuiOs/poupeme-server/src/middlewares"
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
	routes = append(routes, loginRoutes)

	for _, route := range routes {
		if route.NeedAuth {
			router.MethodFunc(route.Method, route.URI, middlewares.Logger(middlewares.Auth(route.Handler)))
		}

		router.MethodFunc(route.Method, route.URI, middlewares.Logger(route.Handler))
	}

	return router
}
