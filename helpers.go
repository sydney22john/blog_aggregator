package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	log.Printf("%s\n", msg)
	type errResp struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errResp{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error: failed marshalling to json: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Write(data)
}

func reverse[T any](s []T) []T {
	for i := 0; i < len(s)/2; i++ {
		temp := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = temp
	}
	return s
}

func contains[T comparable](item T, array []T) bool {
	for _, e := range array {
		if item == e {
			return true
		}
	}
	return false
}
