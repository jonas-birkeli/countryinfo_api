package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// getCountryCodeFromPath separates the country from the path and returns it
func getCountryCodeFromPath(path, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}

// getQueryInt returns the int specified in the query parameter
func getQueryInt(r *http.Request, key string, defaultValue int) int {
	query := r.URL.Query().Get(key)
	if query == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(query)
	if err != nil {
		return defaultValue
	}
	return val
}

// getYearRange takes returns a tuple of the year range specified in the URL
func getYearRange(r *http.Request, key string) (int, int) {
	limitStr := r.URL.Query().Get(key)
	if limitStr == "" {
		return 0, 0
	}

	years := strings.Split(limitStr, "-")
	if len(years) != 2 {
		return 0, 0
	}

	startYear, err1 := strconv.Atoi(years[0])
	endYear, err2 := strconv.Atoi(years[1])

	if err1 != nil || err2 != nil {
		return 0, 0
	}

	return startYear, endYear
}

// getAPIStatus attempts to get a response from the apiURL, and returns its response.
func getAPIStatus(apiURL string) (int, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(apiURL)
	if err != nil {
		return 0, fmt.Errorf("Failed to get API status: %v", err)
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}
