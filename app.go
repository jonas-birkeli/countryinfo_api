package main

import (
	"net/http"
	"time"
)

type App struct {
	startTime time.Time
	client    *http.Client
}
