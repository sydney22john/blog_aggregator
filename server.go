package main

import (
	"log"
	"net/http"
)

func runServer(port string) {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":" + port,
		Handler: middlewareCors(mux),
	}

	mux.HandleFunc("/v1/readiness", handlerReadiness)
	mux.HandleFunc("/v1/err", handlerErr)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
