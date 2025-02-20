package responses

// PopulationResponse represents the response for population data
type PopulationResponse struct {
	Mean   int                   `json:"mean"`
	Values []PopulationYearValue `json:"values"`
}

// PopulationYearValue represents population for a specific year
type PopulationYearValue struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

// isResponse is a marker function to differentiate between response types
func (r PopulationResponse) isResponse() {}

type FieldsName struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
}
