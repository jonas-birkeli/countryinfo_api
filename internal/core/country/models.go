package country

import (
	"context"
)

// CountryInfo represents country information from both APIs
type CountryInfo struct {
	Name       string
	Continents []string
	Population int
	Languages  map[string]string
	Borders    []string
	Flag       string
	Capital    string
	Cities     []string
}

// Service interface defines methods for country operations
type Service interface {
	GetCountryInfo(ctx context.Context, code string, cityLimit string) (*CountryInfo, error)
}
