// config/config.go
package config

import (
	"os"
	"time"
)

// Config holds all configuration values
type Config struct {
	Port         string
	BaseAPI      string
	StartTime    time.Time
	ExternalAPIs ExternalAPIConfig
	Endpoints    EndpointConfig
	ContentType  string
}

type ExternalAPIConfig struct {
	CountriesNowAPI  string
	RestCountriesAPI string
	TestEndpoints    TestEndpointConfig
	Paths            PathConfig
}

type TestEndpointConfig struct {
	CountriesNow  string
	RestCountries string
}

type EndpointConfig struct {
	Info       string
	Population string
	Status     string
}

type PathConfig struct {
	Population string
	Flag       string
	Cities     string
	Capital    string
}

// NewConfig creates a new configuration with default values
func NewConfig() *Config {
	return &Config{
		Port:        getEnvOrDefault("PORT", "8080"),
		BaseAPI:     "/countryinfo/v1",
		StartTime:   time.Now(),
		ContentType: "application/json",
		ExternalAPIs: ExternalAPIConfig{
			CountriesNowAPI:  "http://129.241.150.113:3500/api/v0.1",
			RestCountriesAPI: "http://129.241.150.113:8080/v3.1",
			TestEndpoints: TestEndpointConfig{
				CountriesNow:  "/countries",
				RestCountries: "/all",
			},
			Paths: PathConfig{
				Population: "/countries/population",
				Flag:       "/countries/flag/images",
				Cities:     "/countries/cities",
				Capital:    "/countries/capital",
			},
		},
	}
}

// Init initializes derived configuration values
func (c *Config) Init() {
	c.Endpoints = EndpointConfig{
		Info:       c.BaseAPI + "/info/",
		Population: c.BaseAPI + "/population/",
		Status:     c.BaseAPI + "/status/",
	}
}

// Helper function to get environment variable with default fallback
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
