package handlers

import (
	"assignment_1/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// RequestCountries requests info about all countries from CountriesNowAPI, returns a struct of it.
func RequestCountries() CountryResponse {
	resp, err := http.Get(config.CountriesNowAPI + "/countries")
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var countries CountryResponse

	err = json.Unmarshal(body, &countries)
	if err != nil {
		panic(err)
	}

	return countries
}

func RequestCountryInfo[T any](endpoint, param string, paramValue string) T {
	jsonData := map[string]interface{}{
		param: paramValue,
	}

	jsonPayload, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(
		endpoint,
		config.ApplicationOrJsonSpecifier,
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var info T
	err = json.Unmarshal(body, &info)
	if err != nil {
		panic(err)
	}

	return info
}

var allCountries = RequestCountries().Data

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.ApplicationOrJsonSpecifier)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Allow", http.MethodGet) // Inform a client that only GET is allowed
		_, err := fmt.Fprintln(w, `{"error": true, "msg": "Method not allowed. Use GET."}`)
		if err != nil {
			return
		}
		return
	}

	countryCodeAsIso2 := getCountryCodeFromPath(r.URL.Path, config.EndpointInfo)
	limit := getQueryInt(r, "limit", 10)

	if countryCodeAsIso2 == "" {
		// No country given
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, `{"error": true, "msg": "Specify a country ../info/{no}"}`, http.StatusBadRequest)
		return
	}

	var response InfoResponse

	country := ""
	for _, _country := range allCountries {
		if strings.ToUpper(_country.Iso2) == strings.ToUpper(countryCodeAsIso2) {
			country = _country.Country
		}
	}
	if country == "" {
		// Found no country with Iso2 code.
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, `{"error": true, "msg": "That's a countrycode I have never heard of before.."}`, http.StatusBadRequest)
		return
	}

	flagInfo := RequestCountryInfo[FlagInfoResponse](config.CountriesNowAPI+config.PathFlag, "iso2", countryCodeAsIso2)
	capitalInfo := RequestCountryInfo[CapitalInfoResponse](config.CountriesNowAPI+config.PathCapital, "country", country)
	citiesInfo := RequestCountryInfo[CitiesInfoResponse](config.CountriesNowAPI+config.PathCities, "country", country)
	populationInfo := RequestCountryInfo[PopulationInfoResponse](config.CountriesNowAPI+config.PathPopulation, "country", country)

	response.Name = country
	response.Flag = flagInfo.Data.Flag
	response.Capital = capitalInfo.Data.Capital

	// Try to limit if withing bounds
	if limit < len(citiesInfo.Data) {
		response.Cities = citiesInfo.Data[:limit]
	} else {
		response.Cities = citiesInfo.Data
	}

	// Get the latest population statistics
	n := len(populationInfo.Data.PopulationCounts)
	if n > 0 {
		lastPopulationRecord := populationInfo.Data.PopulationCounts[n-1]
		response.Population = lastPopulationRecord.Value
	}

	w.Header().Set("Content-Type", config.ApplicationOrJsonSpecifier)

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, `{"error": true, "msg": "Failed to encode JSON"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, `{"error": true, "msg": "Failed to write response"}`, http.StatusInternalServerError)
	}
}
