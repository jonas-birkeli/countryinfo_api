package status

import (
	"context"
	"countryinfo/internal/client/countriesnow"
	"countryinfo/internal/client/restcountries"
	"io"
	"net/http"
	"sync"
	"time"
)

// Service defines methods for status operations
type service struct {
	startTime           time.Time
	countriesNowClient  *countriesnow.Client
	restCountriesClient *restcountries.Client
}

// NewService Service interface defines methods for status operations
func NewService(cnClient *countriesnow.Client, rcClient *restcountries.Client) Service {
	return &service{
		startTime:           time.Now(),
		countriesNowClient:  cnClient,
		restCountriesClient: rcClient,
	}
}

// GetStatus returns status information
func (s *service) checkCountriesNowAPI() int {
	resp, err := http.Head(s.countriesNowClient.GetBaseURL() + "/countries")
	if err != nil {
		return -1
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

// GetStatus returns status information
func (s *service) checkRestCountriesAPI() int {
	resp, err := http.Head(s.restCountriesClient.GetBaseURL() + "/alpha/no")
	if err != nil {
		return -1
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	return resp.StatusCode
}

// GetStatus returns status information concurrently
func (s *service) GetStatus(ctx context.Context) (*InfoStatus, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var cnStatus int
	var rcStatus int

	go func() {
		defer wg.Done()
		// Check CountriesNow API
		cnStatus = s.checkCountriesNowAPI()
	}()

	go func() {
		defer wg.Done()
		// Check RestCountries API
		rcStatus = s.checkRestCountriesAPI()
	}()

	wg.Wait()

	return &InfoStatus{
		CountriesNowAPI:  cnStatus,
		RestCountriesAPI: rcStatus,
		Version:          "v1",
		Uptime:           int(time.Since(s.startTime).Seconds()),
	}, nil
}
