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
	periodService := service.NewPeriodService(config.db)
	periodHandler := handlers.NewPeriodHandler(userService, periodService)
	categoryService := service.NewCategoryService(config.db)
	categoryHandler := handlers.NewCategoryHandler(userService, categoryService)
	methodService := service.NewMethodService(config.db)
	methodHandler := handlers.NewMethodHandler(userService, methodService)
	r.Group(func(protected chi.Router) {
		protected.Use(auth.RequireFirebaseUser(config.firebaseAuth))
		protected.Get("/api/v1/auth/me", handlers.Me(userService))
		protected.Post("/api/v1/salary", handlers.AddSalary)
		protected.Post("/api/v1/budget/calculate", handlers.CalculateBudget)
		protected.Post("/api/v1/budget/custom", handlers.CalculateCustomBudget)
		protected.Post("/api/v1/periods", periodHandler.CreateOrGetPeriod)
		protected.Get("/api/v1/periods/{periodID}", periodHandler.GetPeriod)
		protected.Post("/api/v1/periods/{periodID}/financial-summary", periodHandler.SaveFinancialSummary)
		protected.Post("/api/v1/periods/{periodID}/financial-summary/preview", periodHandler.PreviewFinancialSummary)
		protected.Post("/api/v1/fixed-expenses", periodHandler.CreateFixedExpense)
		protected.Get("/api/v1/fixed-expenses", periodHandler.ListFixedExpenses)
		protected.Get("/api/v1/fixed-expenses/history", periodHandler.ListHistoricalFixedExpenses)
		protected.Patch("/api/v1/fixed-expenses/{expenseID}", periodHandler.UpdateFixedExpense)
		protected.Get("/api/v1/categories", categoryHandler.List)
		protected.Post("/api/v1/categories", categoryHandler.Create)
		protected.Patch("/api/v1/categories/{categoryID}", categoryHandler.Update)
		protected.Get("/api/v1/methods", methodHandler.List)
		protected.Post("/api/v1/methods", methodHandler.Create)
	})
	return r
}
