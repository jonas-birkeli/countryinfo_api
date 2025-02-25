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

type FieldsName struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
}
