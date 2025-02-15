package country

import (
	"assignment_1/internal/client/countriesnow"
	"assignment_1/internal/client/restcountries"
	"context"
	"errors"
	"strconv"
)

// Service defines methods for country operations
var defaultService Service

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

// InitService initializes the default country service
func InitService(cnClient *countriesnow.Client, rcClient *restcountries.Client) {
	defaultService = NewService(cnClient, rcClient)
}

// GetService returns the default country service
func GetService() Service {
	return defaultService
}

// GetCountryInfo returns information about a country
func (s *service) GetCountryInfo(ctx context.Context, code string, cityLimit string) (*CountryInfo, error) {
	// Get base country info from REST Countries API
	restCountryInfo, err := s.restCountriesClient.GetCountryByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	// Get cities from CountriesNow API - use the common name
	cities, err := s.countriesNowClient.GetCities(ctx, restCountryInfo.Name.Common)
	if err != nil {
		return nil, err
	}

	// Parse city limit with default of 10
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
