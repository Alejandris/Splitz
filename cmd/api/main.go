package main

import (
	"log"
	"net/http"
	"os"

	"splitz/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/health", handlers.HealthCheck)
	r.Post("/api/v1/salary", handlers.AddSalary)
	r.Post("/api/v1/budget/calculate", handlers.CalculateBudget)

	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	log.Printf("API listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
