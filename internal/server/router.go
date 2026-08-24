package server

import (
	"splitz/internal/handlers"

	"github.com/go-chi/chi/v5"
)

// NewRouter construye el router HTTP de la API.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", handlers.HealthCheck)
	r.Post("/api/v1/salary", handlers.AddSalary)
	r.Post("/api/v1/budget/calculate", handlers.CalculateBudget)
	return r
}
