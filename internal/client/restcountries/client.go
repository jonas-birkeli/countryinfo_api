package restcountries

import (
	"context"
	"countryinfo/internal/config"
	"countryinfo/internal/responses"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client is the client for the RestCountries API
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// CountryInfo is the structure for the response from the RestCountries API
type CountryInfo struct {
	Name struct {
		Common     string `json:"common"`
		Official   string `json:"official"`
		NativeName map[string]struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nativeName"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flags      struct {
		PNG string `json:"png"`
		SVG string `json:"svg"`
	} `json:"flags"`
	Capital []string `json:"capital"`
}

// GetBaseURL returns the default base URL
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// NewClient creates a new RestCountries client
func NewClient(cfg *config.Config) *Client {
	return &Client{
		baseURL:    cfg.ExternalAPIs.RestCountriesAPI,
		httpClient: &http.Client{},
	}
}

// GetCountryByCode retrieves country information by country code
func (c *Client) GetCountryByCode(ctx context.Context, code string) (*CountryInfo, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf("%s/alpha/%s", c.baseURL, code),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: status code %d", resp.StatusCode)
	}

	// Decode into an array of CountryInfo since the API returns array
	var countries []CountryInfo
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	// Return error if no countries found
	if len(countries) == 0 {
		return nil, fmt.Errorf("no country found with code: %s", code)
	}

	// Return first (and should be only) country

	return &countries[0], nil
}

// TranslateCountryCode translates the Iso2-code to its common country name
func (c *Client) TranslateCountryCode(ctx context.Context, code string) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf("%s/alpha/%s?fields=name", c.baseURL, code),
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: status code %d", resp.StatusCode)
	}

	var fieldName responses.FieldsName
	if err := json.NewDecoder(resp.Body).Decode(&fieldName); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	return fieldName.Name.Common, nil
}
