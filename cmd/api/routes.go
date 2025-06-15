package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// PublicRoutes
	mux.Get("/health", app.Health)

	//mux.Post("/users/submit", app.Registration)
	//mux.Post("/users/login", app.Login)

	// "For testing purposes, these endpoints are public. They will become private later."
	mux.Delete("/users", app.DeleteUser)
	mux.Delete("/vulnerable/users", app.DeleteUserSQLi)

	// PrivateRoutes
	mux.Route("/admin", func(mux chi.Router) {

		//CRUD OPERATIONS
		//mux.Get("/users/get/{id}", app.GetUser)
		//mux.Put("/users/get/{id}", app.EditUser)
		//mux.Get("/users/", app.GetAllUsers)

	})

	return mux
}
