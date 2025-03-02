package population

import (
	"context"
	"countryinfo/internal/client/countriesnow"
	"countryinfo/internal/client/restcountries"
	"errors"
)

// Service defines methods for population operations
type service struct {
	countriesNowClient  *countriesnow.Client
	restCountriesClient *restcountries.Client
}

// NewService Service interface defines methods for population operations
func NewService(cnClient *countriesnow.Client, rcClient *restcountries.Client) Service {
	return &service{
		countriesNowClient:  cnClient,
		restCountriesClient: rcClient,
	}
}

// GetPopulationData returns population data for a country.
// Attempts with both official and common name.
func (s *service) GetPopulationData(ctx context.Context, code string, timeRange *TimeRange) (*PopulationData, error) {
	commonName, officialName, err := s.restCountriesClient.TranslateCountryCode(ctx, code)
	if err != nil {
		return nil, err
	}

	// Get population data from CountriesNow API
	data, err := s.countriesNowClient.GetPopulation(ctx, commonName)
	if err != nil {
		data, err = s.countriesNowClient.GetPopulation(ctx, officialName)
		if err != nil {
			return nil, err
		}
	}

	// Filter by time range if provided
	var filteredValues []YearValue
	var sum int
	count := 0

	for _, v := range data {
		if timeRange != nil {
			if v.Year < timeRange.StartYear || v.Year > timeRange.EndYear {
				continue
			}
		}

		filteredValues = append(filteredValues, YearValue{
			Year:  v.Year,
			Value: v.Value,
		})
		sum += v.Value
		count++
	}

	if count == 0 {
		return nil, errors.New("no population data found for given time range")
	}

	return &PopulationData{
		Mean:   sum / count,
		Values: filteredValues,
	}, nil
}
