package middlewares

import (
	"log"
	"net/http"

	"github.com/zGuiOs/poupeme-server/src/auth"
	"github.com/zGuiOs/poupeme-server/src/responses"
)

// TODO: make something util with this
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Auth verify if the user is authorized
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Erro(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
