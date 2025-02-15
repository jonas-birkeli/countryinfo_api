// internal/core/status/service.go
package status

import (
	"assignment_1/internal/client/countriesnow"
	"assignment_1/internal/client/restcountries"
	"context"
	"net/http"
	"time"
)

type service struct {
	startTime           time.Time
	countriesNowClient  *countriesnow.Client
	restCountriesClient *restcountries.Client
}

func NewService(cnClient *countriesnow.Client, rcClient *restcountries.Client) Service {
	return &service{
		startTime:           time.Now(),
		countriesNowClient:  cnClient,
		restCountriesClient: rcClient,
	}
}



func (s *service) checkCountriesNowAPI() string {
	resp, err := http.Get(s.countriesNowClient.GetBaseURL() + "/countries")
	if err != nil {
		return "Error"
	}
	defer resp.Body.Close()
	return http.StatusText(resp.StatusCode)
}

func (s *service) checkRestCountriesAPI() string {
	resp, err := http.Get(s.restCountriesClient.GetBaseURL() + "/all")
	if err != nil {
		return "Error"
	}
	defer resp.Body.Close()
	return http.StatusText(resp.StatusCode)
}

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