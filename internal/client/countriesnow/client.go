package countriesnow

import (
	"bytes"
	"context"
	"countryinfo/internal/config"
	"countryinfo/internal/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client is the client for the CountriesNow API
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// GetBaseURL returns the base URL
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// NewClient creates a new CountriesNow client
func NewClient(cfg *config.Config) *Client {
	client := &Client{
		baseURL:    cfg.ExternalAPIs.CountriesNowAPI,
		httpClient: &http.Client{},
	}
	return client
}

// cityRequest is the structure for requesting cities
type cityRequest struct {
	Limit   int    `json:"limit"`
	Order   string `json:"order"`
	OrderBy string `json:"orderBy"`
	Country string `json:"country"`
}

// GetCities retrieves cities for a country
func (c *Client) GetCities(ctx context.Context, country string) ([]string, error) {
	reqBody := cityRequest{
		Country: country,
	}
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

	var response responses.CountriesNowResponse
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

// populationEntry represents population data for a specific year
type populationEntry struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

// YearValue represents population for a specific year
type YearValue struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

// GetPopulation retrieves population data for a country
func (c *Client) GetPopulation(ctx context.Context, countryName string) ([]YearValue, error) {
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

	var response responses.CountriesNowResponse
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
func (c *Client) doRequest(req *http.Request, response *responses.CountriesNowResponse) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	if response.Error {
		return fmt.Errorf("API error: %s", response.Msg)
	}

	return nil
}
