package handlers

import (
	"encoding/json"
	"net/http"
)

// writeJSONResponse writes a JSON response to the http.ResponseWriter
func writeJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// validateCountryCode validates the two-letter country code
func validateCountryCode(code string) bool {
	return len(code) == 2
}
