package status

import (
	"context"
	"countryinfo/internal/client/countriesnow"
	"countryinfo/internal/client/restcountries"
	"net/http"
	"time"
)

// Service defines methods for status operations
type service struct {
	startTime           time.Time
	countriesNowClient  *countriesnow.Client
	restCountriesClient *restcountries.Client
}

// Service interface defines methods for status operations
func NewService(cnClient *countriesnow.Client, rcClient *restcountries.Client) Service {
	return &service{
		startTime:           time.Now(),
		countriesNowClient:  cnClient,
		restCountriesClient: rcClient,
	}
}

// GetStatus returns status information
func (s *service) checkCountriesNowAPI() string {
	resp, err := http.Get(s.countriesNowClient.GetBaseURL() + "/countries")
	if err != nil {
		return "Error"
	}
	defer resp.Body.Close()
	return http.StatusText(resp.StatusCode)
}

// GetStatus returns status information
func (s *service) checkRestCountriesAPI() string {
	resp, err := http.Get(s.restCountriesClient.GetBaseURL() + "/all")
	if err != nil {
		return "Error"
	}
	defer resp.Body.Close()
	return http.StatusText(resp.StatusCode)
}

// GetStatus returns status information
func (s *service) GetStatus(ctx context.Context) (*StatusInfo, error) {
	// Check CountriesNow API
	cnStatus := s.checkCountriesNowAPI()

	// Check RestCountries API
	rcStatus := s.checkRestCountriesAPI()

	return &StatusInfo{
		CountriesNowAPI:  cnStatus,
		RestCountriesAPI: rcStatus,
		Version:          "v1",
		Uptime:           int64(time.Since(s.startTime).Seconds()),
	}, nil
}
