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
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Access countriesNowAPI
	countriesNowAPIStatus, err := getAPIStatus(config.COUNTRIES_NOW_API_ENDPOINT)
	if err != nil {
		return
	}

	// Access restCountriesAPI
	restCountriesAPIStatus, err := getAPIStatus(config.REST_COUNTRIES_API_ENDPOINT)
	if err != nil {
		return
	}

	uptime := time.Since(startTime).Seconds()

	status := StatusResponse{
		CountriesNowAPI:  countriesNowAPIStatus,
		RestCountriesAPI: restCountriesAPIStatus,
		Version:          config.VERSION,
		Uptime:           uptime,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %s", err), http.StatusInternalServerError)
		return
	}
}
