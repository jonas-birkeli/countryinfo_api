package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const LINEBREAK = "\n"

func DiagHandler(w http.ResponseWriter, r *http.Request) {

	output := "Request:" + LINEBREAK
	output += "URL Path " + r.URL.Path + LINEBREAK
	output += "Method " + r.Method + LINEBREAK

	output += LINEBREAK + "Headers:" + LINEBREAK
	for k, v := range r.Header {
		for _, vv := range v {
			output += k + ": " + vv + LINEBREAK
		}
	}

	// Body
	content, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error during body extraction: " + err.Error())
		http.Error(w, "Error during decoding", http.StatusInternalServerError)
		return
	}

	output += LINEBREAK + "Content: " + LINEBREAK
	output += string(content)

	_, err = fmt.Fprintf(w, "%v", output)
	if err != nil {
		log.Print("An error occured: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
