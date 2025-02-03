package handlers

import (
	"assignment_1/config"
	"fmt"
	"net/http"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	countryCode := getCountryCodeFromPath(r.URL.Path, config.INFO_ENDPOINT)
	limit := getQueryInt(r, "limit", 10)

	_, err := fmt.Fprintf(w, "Country: %s, Limit: %d\n", countryCode, limit)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
