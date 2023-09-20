package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {

	payloadData, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal json: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payloadData)
}

func responseWithError(w http.ResponseWriter, code int, message string) {

	if code > 499 {
		log.Printf("Error 5xx: %v", message)
	}

	type ErrResponse struct {
		Error string `json:"error"`
	}

	responseWithJson(w, code, ErrResponse{Error: message})
}