package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/dayterr/gophkeeper_diploma/internal/handlers"
)

func CreateRouterWithAsyncHandler(ah *handlers.AsyncHandler) chi.Router {
	r := chi.NewRouter()

	r.Use(ah.AuthMiddleware)
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", ah.RegisterUser)
		r.Post("/login", ah.LogUser)
	})

	r.Route("/{dataType}", func(r chi.Router) {
		r.Post("/", ah.PostData)
		r.Get("/", ah.ListData)
		r.Delete("/{dataID}", ah.DeleteData)
	})

	return r
}
