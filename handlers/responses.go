package handlers

type InfoResponse struct {
	/*
		{
			"name": "Norway",
			"continents": ["Europe"],
			"population": 4700000,
			"languages": {"nno":"Norwegian Nynorsk","nob":"Norwegian Bokmål","smi":"Sami"},
			"borders": ["FIN","SWE","RUS"],
			"flag": "https://flagcdn.com/w320/no.png",
			"capital": "Oslo",
			"cities": ["Abelvaer","Adalsbruk","Adland","Agotnes","Agskardet..."]
		}
	*/
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities"`
}

type CountryResponse struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  []struct {
		Iso2    string `json:"Iso2"`
		Iso3    string `json:"Iso3"`
		Country string `json:"country"`
	} `json:"data"`
}

type FlagInfoResponse struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		Name string `json:"name"`
		Flag string `json:"flag"`
		Iso2 string `json:"iso2"`
		Iso3 string `json:"iso3"`
	} `json:"data"`
}

type CapitalInfoResponse struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		Name    string `json:"name"`
		Capital string `json:"capital"`
		Iso2    string `json:"iso2"`
		Iso3    string `json:"iso3"`
	} `json:"data"`
}

type CitiesInfoResponse struct {
	Error bool     `json:"error"`
	Msg   string   `json:"msg"`
	Data  []string `json:"data"`
}

type PopulationInfoResponse struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		Country          string `json:"country"`
		Code             string `json:"code"`
		Iso3             string `json:"iso3"`
		PopulationCounts []struct {
			Year  int `json:"year"`
			Value int `json:"value"`
		} `json:"populationCounts"`
	} `json:"data"`
}
