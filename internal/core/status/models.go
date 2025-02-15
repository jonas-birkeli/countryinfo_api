// internal/core/status/models.go
package status

import "context"

type Service interface {
	GetStatus(ctx context.Context) (*StatusInfo, error)
}

type StatusInfo struct {
	CountriesNowAPI  string `json:"countriesnowapi"`
	RestCountriesAPI string `json:"restcountriesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}
