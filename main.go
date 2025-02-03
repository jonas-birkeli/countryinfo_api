package main

import (
	"assignment_1/config"
	"assignment_1/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	router := http.NewServeMux()

	port := config.PORT
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	router.HandleFunc("/diag", handlers.DiagHandler)
	router.HandleFunc(config.INFO_ENDPOINT, handlers.InfoHandler)
	router.HandleFunc(config.POPULATION_ENDPOINT, handlers.PopulationHandler)
	router.HandleFunc(config.STATUS_ENDPOINT, handlers.StatusHandler)

	log.Printf("Startinvg server now lol on port %v.", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
