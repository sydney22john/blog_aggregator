package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func runServer(cfg apiConfig, port string) {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":" + port,
		Handler: middlewareCors(mux),
	}

	mux.HandleFunc("GET /v1/readiness", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)
	mux.HandleFunc("POST /v1/users", cfg.handlerPostUsers)
	mux.HandleFunc("GET /v1/users", cfg.handlerGetUsersByApikey)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
