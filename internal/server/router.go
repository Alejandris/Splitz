package server

import (
	"splitz/internal/auth"
	"splitz/internal/handlers"
	"splitz/internal/service"

	firebaseauth "firebase.google.com/go/v4/auth"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type routerConfig struct {
	db           *pgxpool.Pool
	firebaseAuth *firebaseauth.Client
}

type RouterOption func(*routerConfig)

func WithDependencies(db *pgxpool.Pool, firebaseAuth *firebaseauth.Client) RouterOption {
	return func(config *routerConfig) {
		config.db = db
		config.firebaseAuth = firebaseAuth
	}
}

// NewRouter construye el router HTTP de la API.
func NewRouter(options ...RouterOption) *chi.Mux {
	config := routerConfig{}
	for _, option := range options {
		option(&config)
	}

	r := chi.NewRouter()
	r.Get("/health", handlers.HealthCheck)

	if config.db == nil || config.firebaseAuth == nil {
		return r
	}

	userService := service.NewUserService(config.db)
	r.Group(func(protected chi.Router) {
		protected.Use(auth.RequireFirebaseUser(config.firebaseAuth))
		protected.Get("/api/v1/auth/me", handlers.Me(userService))
		protected.Post("/api/v1/salary", handlers.AddSalary)
		protected.Post("/api/v1/budget/calculate", handlers.CalculateBudget)
		protected.Post("/api/v1/budget/custom", handlers.CalculateCustomBudget)
	})
	return r
}
