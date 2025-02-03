package handlers

import "net/http"

func HeiHandler(w http.ResponseWriter, r *http.Request) {
	output := "<!DOCTYPE html><html><body><h1>HTML title</h1><p>Woke shit.</p></body></html>"
	_, err := w.Write([]byte(output))
	if err != nil {
		return
	}
}
