package handlers

import "C"
import (
	"countryinfo/internal/core/country"
	"countryinfo/internal/responses"
	"net/http"
	"strings"
)

// countrySvc is the population service
var countrySvc country.Service

// InitCountryService initializes the population service
func InitCountryService(svc country.Service) {
	countrySvc = svc
}

// InfoHandler handles requests for country information
func InfoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSONResponse(w, http.StatusMethodNotAllowed, responses.ErrorResponse{
			Error: "method not allowed",
		})
		return
	}

	// Extract country code from a path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		writeJSONResponse(w, http.StatusBadRequest, responses.ErrorResponse{
			Error: "invalid path",
		})
		return
	}

	countryCode := parts[4]
	if !validateCountryCode(countryCode) {
		writeJSONResponse(w, http.StatusBadRequest, responses.ErrorResponse{
			Error: "invalid country code",
		})
		return
	}

	// Get query parameters
	limit := r.URL.Query().Get("limit")

	// Call service layer
	info, err := countrySvc.GetCountryInfo(r.Context(), countryCode, limit)
	if err != nil {

		writeJSONResponse(w, http.StatusInternalServerError, responses.ErrorResponse{
			Error: "Failed to get country information",
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, info)
}
