package routes

import (
	"mustaqel/handlers/api/role"
	"mustaqel/handlers/api/user"

	"github.com/go-chi/chi/v5"
)

func SetupApiRoutes() *chi.Mux {

	routes := chi.NewRouter()

	routes.Group(func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", user.GetUsers)
			r.Post("/", user.CreateUser)
			r.Get("/{id}", user.GetUser)
			r.Put("/{id}", user.UpdateUser)
			r.Delete("/{id}", user.DeleteUser)
		})

		r.Route("/roles", func(r chi.Router) {
			r.Get("/", role.GetRoles)
			r.Post("/", role.CreateRole)
			r.Get("/{id}", role.GetRole)
			r.Put("/{id}", role.UpdateRole)
			r.Delete("/{id}", role.DeleteRole)
		})

	})

	return routes
}
