package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"user-service/internal/config"
)

func main() {
	// Charger la configuration
	dbConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		log.Fatalf("Error loading database config: %v", err)
	}

	// Créer un contexte avec annulation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialiser la connexion PostgreSQL
	pgConn, err := config.NewPostgresConnection(ctx, dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}
	defer pgConn.Close(ctx)

	// Gérer les signaux d'arrêt
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Attendre un signal d'arrêt
	<-sigChan
	log.Println("Shutting down...")
}
