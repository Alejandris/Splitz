package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"splitz/internal/auth"
	"splitz/internal/config"
	"splitz/internal/database"
	"splitz/internal/server"
)

func main() {
	ctx := context.Background()
	appConfig, err := config.Load()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	dbPool, err := database.NewPostgresPool(ctx, appConfig.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}
	defer dbPool.Close()

	firebaseAuth, err := auth.NewFirebaseAuth(ctx, appConfig.FirebaseProjectID, appConfig.FirebaseCredentialsFile)
	if err != nil {
		log.Fatalf("Firebase connection error: %v", err)
	}

	r := server.NewRouter(server.WithDependencies(dbPool, firebaseAuth))

	addr := ":" + appConfig.Port

	log.Printf("API listening on %s", addr)
	srv := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
