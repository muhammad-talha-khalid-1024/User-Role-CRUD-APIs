package routes

import (
	"mustaqel/handlers/web/index"

	"github.com/go-chi/chi/v5"
)

func SetupWebRoutes() *chi.Mux {
	routes := chi.NewRouter()
	routes.Get("/", index.HomeHandler)
	return routes
}
