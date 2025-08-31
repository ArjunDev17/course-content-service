package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ArjunDev17/course-content-service/config"
	"github.com/ArjunDev17/course-content-service/pkg/db"
	"github.com/ArjunDev17/course-content-service/server/api"
)

func main() {
	// Load config
	configPath := "config/config.yaml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// fallback to environment override
		configPath = ""
	}
	if configPath != "" {
		config.LoadConfig(configPath)
	}

	// connect mongo
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := db.ConnectMongo(ctx)
	if err != nil {
		log.Fatalf("mongo connect error: %v", err)
	}
	// set global client for repository helper
	db.Client = client

	// start server
	router := api.NewRouter()
	port := fmt.Sprintf(":%d", config.Cfg.Server.Port)
	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// graceful shutdown
	go func() {
		log.Printf("starting server at %s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()
	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	// close mongo
	if err := client.Disconnect(ctxShutdown); err != nil {
		log.Printf("mongo disconnect error: %v", err)
	}
	log.Println("Server exiting")
}
