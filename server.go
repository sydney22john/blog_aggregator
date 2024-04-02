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

	mux.HandleFunc("/v1/readiness", handlerReadiness)
	mux.HandleFunc("/v1/err", handlerErr)
	mux.HandleFunc("/v1/users", cfg.handlerPostUsers)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
