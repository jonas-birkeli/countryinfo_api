package handlers

import (
	"encoding/json"
	"net/http"
	"unicode"
)

// writeJSONResponse writes a JSON response to the http.ResponseWriter
func writeJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// validateCountryCode validates the two-letter country code
func validateCountryCode(code string) bool {
	if len(code) != 2 {
		return false
	}

	for _, char := range code {
		if !unicode.IsLetter(char) {
			return false
		}
	}

	return true
}
