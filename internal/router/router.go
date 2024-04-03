package router

import (
	"net/http"

	"github.com/pwnsbd/footsteps/internal/handlers"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlers.HomeHandler(w, r)
	case "/contact":
		handlers.ContactHandler(w, r)
	case "/projects":
		handlers.ProjectHandler(w, r)
	default:
		// Page not found
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}
