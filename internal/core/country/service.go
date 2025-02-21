package country

import (
	"context"
	"countryinfo/internal/client/countriesnow"
	"countryinfo/internal/client/restcountries"
	"errors"
	"strconv"
)

// service implements the Service interface
type service struct {
	countriesNowClient  *countriesnow.Client
	restCountriesClient *restcountries.Client
}

// NewService creates a new country service
func NewService(cnClient *countriesnow.Client, rcClient *restcountries.Client) Service {
	return &service{
		countriesNowClient:  cnClient,
		restCountriesClient: rcClient,
	}
}

// GetCountryInfo returns information about a country
func (s *service) GetCountryInfo(ctx context.Context, code string, cityLimit string) (*CountryInfo, error) {

	// Get base country info from REST Countries API
	restCountryInfo, err := s.restCountriesClient.GetCountryByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	// Parse the city limit with default of 10
	limit := 10 // Default limit
	if cityLimit != "" {
		parsedLimit, err := strconv.Atoi(cityLimit)
		if err != nil {
			return nil, errors.New("invalid city limit")
		}
		if parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	// Get cities from CountriesNow API - use the common name
	cities, err := s.countriesNowClient.GetCities(ctx, restCountryInfo.Name.Common, limit)
	if err != nil {
		return nil, err
	}

	// Apply the limit
	if len(cities) > limit {
		cities = cities[:limit]
	}

	// Get first capital if there is multiple
	var capital string
	if len(restCountryInfo.Capital) > 0 {
		capital = restCountryInfo.Capital[0]
	}

	return &CountryInfo{
		Name:       restCountryInfo.Name.Common, // Using common name
		Continents: restCountryInfo.Continents,
		Population: restCountryInfo.Population,
		Languages:  restCountryInfo.Languages,
		Borders:    restCountryInfo.Borders,
		Flag:       restCountryInfo.Flags.PNG, // Using PNG URL from Flags struct
		Capital:    capital,                   // Using first capital
		Cities:     cities,
	}, nil
}
