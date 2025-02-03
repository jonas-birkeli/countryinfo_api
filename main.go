package main

import (
	"assignment_1/handlers"
	"log"
	"net/http"
	"os"
)

const ENDPOINT = "/countryinfo"
const VERSION = "/v1"
const INFO_ENDPOINT = ENDPOINT + VERSION + "/info/"
const POPULATION_ENDPOINT = ENDPOINT + VERSION + "/population/"
const STATUS_ENDPOINT = ENDPOINT + VERSION + "/status/"
const CONSTRAINT = "/{countryCode}"
const PORT = "8080"

func main() {

	router := http.NewServeMux()

	subPort := PORT
	if os.Getenv("PORT") != "" {
		subPort = os.Getenv("PORT")
	}
	router.HandleFunc("/diag", handlers.DiagHandler)
	router.HandleFunc(INFO_ENDPOINT, handlers.InfoHandler)
	router.HandleFunc(POPULATION_ENDPOINT, handlers.PopulationHandler)
	router.HandleFunc(STATUS_ENDPOINT, handlers.StatusHandler)

	log.Printf("Startinvg server now lol on port %v.", subPort)
	log.Fatal(http.ListenAndServe(":"+subPort, router))
}
