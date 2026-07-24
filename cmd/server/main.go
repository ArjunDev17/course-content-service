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
	httpHandler "github.com/ArjunDev17/course-content-service/handler/http"
	"github.com/ArjunDev17/course-content-service/pkg/db"
	mongoRepo "github.com/ArjunDev17/course-content-service/repository/mongo"
	"github.com/ArjunDev17/course-content-service/server/api"
	courseService "github.com/ArjunDev17/course-content-service/service/course"
)

func main() {

	// ----------------------------
	// Load Configuration
	// ----------------------------
	configPath := "config/config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configPath = ""
	}

	if configPath != "" {
		config.LoadConfig(configPath)
	}

	// ----------------------------
	// Connect MongoDB
	// ----------------------------
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := db.ConnectMongo(ctx)
	if err != nil {
		log.Fatalf("failed to connect mongo: %v", err)
	}

	db.Client = client

	// ----------------------------
	// Dependency Injection
	// ----------------------------

	// Repository
	courseRepo := mongoRepo.NewCourseRepository()

	// Service
	courseSvc := courseService.NewCourseService(courseRepo)

	// Handlers
	courseHandler := httpHandler.NewCourseHandler(courseSvc)
	healthHandler := httpHandler.NewHealthHandler()

	// ----------------------------
	// Router
	// ----------------------------
	router := api.NewRouter(
		courseHandler,
		healthHandler,
	)

	port := fmt.Sprintf(":%d", config.Cfg.Server.Port)

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// ----------------------------
	// Start Server
	// ----------------------------
	go func() {
		log.Printf("Server started on %s", port)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			log.Fatalf("server error: %v", err)
		}
	}()

	// ----------------------------
	// Graceful Shutdown
	// ----------------------------
	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	log.Println("Shutting down server...")

	ctxShutdown, cancelShutdown := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancelShutdown()

	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	if err := client.Disconnect(ctxShutdown); err != nil {
		log.Printf("Mongo disconnect error: %v", err)
	}

	log.Println("Server exited successfully")
}