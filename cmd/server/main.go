package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ArjunDev17/course-content-service/internal/config"
	"github.com/ArjunDev17/course-content-service/internal/database"
)

func main() {

	// Load application configuration
	cfg := config.Load()

	// Initialize PostgreSQL
	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatalf("failed to initialize postgres: %v", err)
	}
	defer db.Close()

	log.Println("✅ PostgreSQL connected successfully")

	// Temporary HTTP server
	server := &http.Server{
		Addr: ":" + cfg.App.Port,
	}

	// Start server in separate goroutine
	go func() {
		log.Printf("🚀 Server started on port %s", cfg.App.Port)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	// Wait for shutdown signal
	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	log.Println("Shutdown signal received...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}

	log.Println("Application stopped.")
}