package handlers

import (
	"countryinfo/internal/config"
	"countryinfo/internal/core/population"
	"net/http"
)

// RegisterHandlers registers all API handlers
func RegisterHandlers(mux *http.ServeMux, cfg *config.Config, popService population.Service) {
	// Initialize services for handlers
	InitPopulationService(popService)

	// Register routes
	mux.HandleFunc(cfg.Endpoints.Info, InfoHandler)
	mux.HandleFunc(cfg.Endpoints.Population, PopulationHandler)
	mux.HandleFunc(cfg.Endpoints.Status, StatusHandler)
}
