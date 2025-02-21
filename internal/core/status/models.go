package status

import "context"

// Service represents the status service
type Service interface {
	GetStatus(ctx context.Context) (*InfoStatus, error)
}

// InfoStatus represents status information
type InfoStatus struct {
	CountriesNowAPI  string `json:"countriesnowapi"`
	RestCountriesAPI string `json:"restcountriesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}
