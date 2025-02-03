package handlers

import (
	"fmt"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	countryCode := getCountryCodeFromPath(r.URL.Path, "/countryinfo/v1/status/")

	_, err := fmt.Fprintf(w, "Country: %s\n", countryCode)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
