// main.go
package main

import (
	"context"
	"countryinfo/internal/api/handlers"
	"countryinfo/internal/api/middleware"
	"countryinfo/internal/client/countriesnow"
	"countryinfo/internal/client/restcountries"
	"countryinfo/internal/config"
	"countryinfo/internal/core/country"
	"countryinfo/internal/core/population"
	"countryinfo/internal/core/status"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize configuration
	cfg := config.NewConfig()
	cfg.Init()

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Create clients
	countriesNowClient, err := countriesnow.NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize CountriesNow client: %v\n", err)
	}
	restCountriesClient := restcountries.NewClient(cfg)

	// Initialize the service
	country.InitService(countriesNowClient, restCountriesClient)

	// Create services
	populationService := population.NewService(countriesNowClient, restCountriesClient)
	statusService := status.NewService(countriesNowClient, restCountriesClient)

	// Initialize services
	handlers.InitStatusService(statusService)

	// Register handlers with config
	handlers.RegisterHandlers(mux, cfg, populationService)

	// Apply a middleware chain
	handler := middleware.Chain(mux,
		middleware.Recover,
		middleware.Logger,
		middleware.RequestID,
		middleware.CORS,
	)

	// Create server
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel for server errors
	serverErrors := make(chan error, 1)

	// Start server
	go func() {
		log.Printf("Server is starting on port %s", cfg.Port)
		serverErrors <- server.ListenAndServe()
	}()

	// Channel for shutdown signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Wait for shutdown or error
	select {
	case err := <-serverErrors:
		log.Fatalf("Error starting server: %v", err)

	case sig := <-shutdown:
		log.Printf("Start shutdown... Signal: %v", sig)

		// Create shutdown context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Attempt a graceful shutdown
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Graceful shutdown failed: %v", err)
			server.Close()
		}
	}

	log.Println("Server shutdown complete")
}
