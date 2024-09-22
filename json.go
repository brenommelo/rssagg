package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Respond with 5XX error code: %v", msg)
	}

	type errRespose struct {
		Error string `json:"error"` //formata o json em um formato ideal para error
	}

	respondWithJSON(w, code, errRespose{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload) //essa função vai empacotar o que passar no payload para bytes

	if err != nil {
		log.Println("Error to marshal JSON response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
