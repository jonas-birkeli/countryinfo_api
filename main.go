package main

import (
	"context"
	"countryinfo/internal/api/handlers"
	"countryinfo/internal/api/middleware"
	"countryinfo/internal/client/countriesnow"
	"countryinfo/internal/client/restcountries"
	"countryinfo/internal/config"
	"countryinfo/internal/core/info"
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
	countriesNowClient := countriesnow.NewClient(cfg)
	restCountriesClient := restcountries.NewClient(cfg)

	// Create services
	countryService := info.NewService(countriesNowClient, restCountriesClient)
	populationService := population.NewService(countriesNowClient, restCountriesClient)
	statusService := status.NewService(countriesNowClient, restCountriesClient)

	// Initialize services (making them global)
	handlers.InitCountryService(countryService)
	handlers.InitPopulationService(populationService)
	handlers.InitStatusService(statusService)

	// Register handlers with config
	handlers.RegisterHandlers(mux, cfg)

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
		ReadTimeout:  cfg.Timeout.ReadTimeout * time.Second,
		WriteTimeout: cfg.Timeout.WriteTimeout * time.Second,
		IdleTimeout:  cfg.Timeout.IdleTimeout * time.Second,
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
