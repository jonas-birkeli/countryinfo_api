package handlers

import (
	"assignment_1/config"
	"fmt"
	"net/http"
)

func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	countryCode := getCountryCodeFromPath(r.URL.Path, config.PopulationEndpoint)
	startYear, endYear := getYearRange(r, "limit")

	_, err := fmt.Fprintf(w, "Country: %s, Startyear: %d, Endyear: %d\n", countryCode, startYear, endYear)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

}
