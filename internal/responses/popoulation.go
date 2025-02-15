// population.go
package responses

// PopulationResponse represents the response for population data
type PopulationResponse struct {
	Mean   int                   `json:"mean"`
	Values []PopulationYearValue `json:"values"`
}

type PopulationYearValue struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

func (r PopulationResponse) isResponse() {}

// StatusResponse represents the status response
type StatusResponse struct {
	CountriesNowAPI  string `json:"countriesnowapi"`
	RestCountriesAPI string `json:"restcountriesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}

func (r StatusResponse) isResponse() {}
