package status

import "context"

// Service represents the status service
type Service interface {
	GetStatus(ctx context.Context) (*StatusInfo, error)
}

// StatusInfo represents status information
type StatusInfo struct {
	CountriesNowAPI  string `json:"countriesnowapi"`
	RestCountriesAPI string `json:"restcountriesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}
