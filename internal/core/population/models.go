package population

import (
	"context"
)

// PopulationData represents population data for a country
type PopulationData struct {
	Mean   int
	Values []YearValue
}

// YearValue represents population for a specific year
type YearValue struct {
	Year  int
	Value int
}

// TimeRange represents a range of years
type TimeRange struct {
	StartYear int
	EndYear   int
}

// Service interface defines methods for population operations
type Service interface {
	GetPopulationData(ctx context.Context, code string, timeRange *TimeRange) (*PopulationData, error)
}
