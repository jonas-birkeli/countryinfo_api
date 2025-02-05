package config

const (
	Port    = "8080"
	BaseAPI = "/countryinfo/v1"

	// Local endpoints
	EndpointInfo       = BaseAPI + "/info/"
	EndpointPopulation = BaseAPI + "/population/"
	EndpointStatus     = BaseAPI + "/status/"

	// External APIs base URLs
	CountriesNowAPI                              = "http://129.241.150.113:3500/api/v0.1"
	RestCountriesAPI                             = "http://129.241.150.113:8080/v3.1"
	CountriesNowAPIAccessableEndpointForTesting  = "/countries"
	RestCountriesAPIAccessableEndpointForTesting = "/all"

	// Path suffixes
	PathPopulation = "/countries/population"
	PathFlag       = "/countries/flag/images"
	PathCities     = "/countries/cities"
	PathCapital    = "/countries/capital"

	ApplicationOrJsonSpecifier = "application/json"
)
