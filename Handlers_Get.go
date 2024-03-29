package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct {
		Status string `json:"status"`
	}{"ok"})
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "internal server error")
}
