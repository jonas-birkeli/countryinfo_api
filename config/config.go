package config

const PORT = "8080"

const BASE_URL = "/countryinfo"
const VERSION = "/v1"

const TRAILING_SLASH = "/"

const INFO_ENDPOINT = BASE_URL + VERSION + "/info" + TRAILING_SLASH
const POPULATION_ENDPOINT = BASE_URL + VERSION + "/population" + TRAILING_SLASH
const STATUS_ENDPOINT = BASE_URL + VERSION + "/status" + TRAILING_SLASH
