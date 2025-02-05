package handlers

import (
	"assignment_1/config"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type StatusResponse struct {
	CountriesNowAPI  int     `json:"countries_now_api"`
	RestCountriesAPI int     `json:"rest_countries_api"`
	Version          string  `json:"version"`
	Uptime           float64 `json:"uptime"`
}

var startTime = time.Now()

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.ApplicationOrJsonSpecifier)
	if r.Method != http.MethodGet {

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Access countriesNowAPI
	countriesNowAPIStatus, err := getAPIStatus(config.CountriesNowAPI, config.CountriesNowAPIAccessableEndpointForTesting)
	if err != nil {
		return
	}

	// Access restCountriesAPI
	restCountriesAPIStatus, err := getAPIStatus(config.RestCountriesAPI, config.RestCountriesAPIAccessableEndpointForTesting)
	if err != nil {
		return
	}

	uptime := time.Since(startTime).Seconds()

	status := StatusResponse{
		CountriesNowAPI:  countriesNowAPIStatus,
		RestCountriesAPI: restCountriesAPIStatus,
		Version:          "v1",
		Uptime:           uptime,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %s", err), http.StatusInternalServerError)
		return
	}
}
