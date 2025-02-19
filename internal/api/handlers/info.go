package handlers

import "C"
import (
	"countryinfo/internal/core/country"
	"countryinfo/internal/responses"
	"net/http"
	"strings"
)

// InfoHandler handles requests for country information
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONResponse(w, http.StatusMethodNotAllowed, responses.ErrorResponse{
			Error: "Method not allowed",
		})
		return
	}

	// Extract country code from path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		writeJSONResponse(w, http.StatusBadRequest, responses.ErrorResponse{
			Error: "Invalid path",
		})
		return
	}

	countryCode := parts[4]
	if !validateCountryCode(countryCode) {
		writeJSONResponse(w, http.StatusBadRequest, responses.ErrorResponse{
			Error: "Invalid country code",
		})
		return
	}

	// Get query parameters
	limit := r.URL.Query().Get("limit")
	// Add limit parsing logic here

	// Call service layer
	info, err := country.GetService().GetCountryInfo(r.Context(), countryCode, limit)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, responses.ErrorResponse{
			Error: "Failed to get country information",
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, info)
}
