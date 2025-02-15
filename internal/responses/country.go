package responses

// CountryInfoResponse represents the response for country information
type CountryInfoResponse struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities,omitempty"`
}

// CountryInfoResponse represents the response for country information
func (r CountryInfoResponse) isResponse() {}
