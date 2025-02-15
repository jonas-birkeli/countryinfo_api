// client/countriesnow/client.go
package countriesnow

import (
	"assignment_1/internal/config"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	// Add a map for ISO to country name mapping
	isoToCountry map[string]string
	mu           sync.RWMutex // For safe concurrent access
}

type countryListResponse struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  []struct {
		Iso2    string `json:"iso2"`
		Country string `json:"country"`
	} `json:"data"`
}

// GetBaseURL returns the base URL
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

func NewClient(cfg *config.Config) (*Client, error) {
	client := &Client{
		baseURL:      cfg.ExternalAPIs.CountriesNowAPI,
		httpClient:   &http.Client{},
		isoToCountry: make(map[string]string),
	}

	// Initialize the ISO to country mapping
	if err := client.initializeCountryMap(); err != nil {
		return nil, fmt.Errorf("failed to initialize country mapping: %w", err)
	}

	return client, nil
}

func (c *Client) initializeCountryMap() error {
	req, err := http.NewRequest("GET", c.baseURL+"/countries", nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var response countryListResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	c.mu.Lock()
	for _, country := range response.Data {
		c.isoToCountry[country.Iso2] = country.Country
	}
	c.mu.Unlock()

	return nil
}

// GetCountryName returns the full country name for an ISO code
func (c *Client) GetCountryName(isoCode string) (string, error) {
	c.mu.RLock()
	countryName, ok := c.isoToCountry[isoCode]
	c.mu.RUnlock()

	if !ok {
		return "", fmt.Errorf("no country found for ISO code: %s", isoCode)
	}
	return countryName, nil
}

// countriesnowResponse is a generic response structure
type countriesnowResponse struct {
	Error bool        `json:"error"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

// cityRequest is the structure for requesting cities
type cityRequest struct {
	Country string `json:"country"`
}

// GetCities retrieves cities for a country
func (c *Client) GetCities(ctx context.Context, country string) ([]string, error) {
	reqBody := cityRequest{Country: country}
	jsonBody, err := json.Marshal(reqBody)

	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/countries/cities",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	var response countriesnowResponse
	response.Data = &[]string{}

	if err := c.doRequest(req, &response); err != nil {
		return nil, err
	}

	cities, ok := response.Data.(*[]string)
	if !ok {
		return nil, fmt.Errorf("invalid response data type")
	}

	return *cities, nil
}

// populationResponse represents population data
type populationResponse struct {
	CountryName string            `json:"country"`
	Population  []populationEntry `json:"populationCounts"`
}

type populationEntry struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

type YearValue struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

// GetPopulation retrieves population data for a country
func (c *Client) GetPopulation(ctx context.Context, isoCode string) ([]YearValue, error) {
	countryName, err := c.GetCountryName(isoCode)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/countries/population",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{"country": "%s"}`, countryName))),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	var response countriesnowResponse
	response.Data = &populationResponse{}

	if err := c.doRequest(req, &response); err != nil {
		return nil, err
	}

	popData, ok := response.Data.(*populationResponse)
	if !ok {
		return nil, fmt.Errorf("invalid response data type")
	}

	yearValues := make([]YearValue, len(popData.Population))
	for i, p := range popData.Population {
		yearValues[i] = YearValue{
			Year:  p.Year,
			Value: p.Value,
		}
	}

	return yearValues, nil
}

// doRequest performs the HTTP request and unmarshals the response
func (c *Client) doRequest(req *http.Request, response *countriesnowResponse) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	if response.Error {
		return fmt.Errorf("API error: %s", response.Msg)
	}

	return nil
}
