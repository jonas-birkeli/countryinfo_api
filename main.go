package main

import (
	"assignment_1/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	router := http.NewServeMux()

	PORT := "23456"
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}

	router.HandleFunc("/diag", handlers.DiagHandler)
	router.HandleFunc("/hei", handlers.HeiHandler)

	log.Printf("Startinvg server now lol on port %v.", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
