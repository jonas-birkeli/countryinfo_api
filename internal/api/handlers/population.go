package handlers

import (
	"assignment_1/internal/core/population"
	"assignment_1/internal/responses"
	"net/http"
	"strconv"
	"strings"
)

// populationSvc is the population service
var populationSvc population.Service

// InitPopulationService initializes the population service
func InitPopulationService(svc population.Service) {
	populationSvc = svc
}

// PopulationHandler handles requests for population data
func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONResponse(w, http.StatusMethodNotAllowed, responses.ErrorResponse{
			Error: "Method not allowed",
		})
		return
	}

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

	// Parse time range from query parameters
	var timeRange *population.TimeRange
	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		years := strings.Split(limitParam, "-")
		if len(years) != 2 {
			writeJSONResponse(w, http.StatusBadRequest, responses.ErrorResponse{
				Error: "Invalid time range format. Expected format: startYear-endYear",
			})
			return
		}

		startYear, err := strconv.Atoi(years[0])
		if err != nil {
			writeJSONResponse(w, http.StatusBadRequest, responses.ErrorResponse{
				Error: "Invalid start year",
			})
			return
		}

		endYear, err := strconv.Atoi(years[1])
		if err != nil {
			writeJSONResponse(w, http.StatusBadRequest, responses.ErrorResponse{
				Error: "Invalid end year",
			})
			return
		}

		timeRange = &population.TimeRange{
			StartYear: startYear,
			EndYear:   endYear,
		}
	}

	// Call service with parsed TimeRange struct
	popData, err := populationSvc.GetPopulationData(r.Context(), countryCode, timeRange)
	if err != nil {

		writeJSONResponse(w, http.StatusInternalServerError, responses.ErrorResponse{
			Error: "Failed to get population data",
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, popData)
}
