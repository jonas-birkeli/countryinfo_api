package config

// This projects constant pathing \\

const PORT = "8080"
const BaseUrl = "/countryinfo/"
const VERSION = "v1"

const TrailingSlash = "/"
const InfoEndpoint = BaseUrl + VERSION + "/info/"
const PopulationEndpoint = BaseUrl + VERSION + "/population/"
const StatusEndpoint = BaseUrl + VERSION + "/status/"

// Constants for APIs used \\

const CountriesNowApiEndpoint = "http://129.241.150.113:3500/api/v0.1/countries/"
const RestCountriesApiEndpoint = "http://129.241.150.113:8080/v3.1/all"

const CountriesNowApiPopulationPath = "population"
const CountriesNowApiFlagPath = "flag/images"
const CountriesNowApiCitiesPath = "cities"
const CountriesNowApiCapitalPath = "capital"
