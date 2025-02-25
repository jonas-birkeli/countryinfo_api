package handlers

import (
	"countryinfo/internal/config"
	"net/http"
)

// RegisterHandlers registers all API handlers
func RegisterHandlers(mux *http.ServeMux, cfg *config.Config) {
	// Register routes
	mux.HandleFunc(cfg.Endpoints.Info, InfoHandler)
	mux.HandleFunc(cfg.Endpoints.Population, PopulationHandler)
	mux.HandleFunc(cfg.Endpoints.Status, StatusHandler)
}
