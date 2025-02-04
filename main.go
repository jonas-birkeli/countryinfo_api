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
	router.HandleFunc(config.InfoEndpoint, handlers.InfoHandler)
	router.HandleFunc(config.PopulationEndpoint, handlers.PopulationHandler)
	router.HandleFunc(config.StatusEndpoint, handlers.StatusHandler)

	log.Printf("Startinvg server now lol on port %v.", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
