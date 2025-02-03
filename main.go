package main

import (
	"assignment_1/handlers"
	"log"
	"net/http"
	"os"
)

const ENDPOINT_INFO = "/countryinfo/v1/info/"
const ENDPOINT_POPULATION = "/countryinfo/v1/population/"
const ENDPOINT_STATUS = "/countryinfo/v1/status/"
const PORT = "8080"

func main() {

	router := http.NewServeMux()

	subPort := PORT
	if os.Getenv("PORT") != "" {
		subPort = os.Getenv("PORT")
	}
	router.HandleFunc(ENDPOINT_INFO, handlers.InfoHandler)
	router.HandleFunc(ENDPOINT_POPULATION, handlers.PopulationHandler)
	router.HandleFunc(ENDPOINT_STATUS, handlers.StatusHandler)

	log.Printf("Startinvg server now lol on port %v.", subPort)
	log.Fatal(http.ListenAndServe(":"+subPort, router))
}
