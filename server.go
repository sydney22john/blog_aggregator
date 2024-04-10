package main

import (
	"log"
	"net/http"
)

func runServer(cfg apiConfig, port string) {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":" + port,
		Handler: middlewareCors(mux),
	}

	mux.HandleFunc("GET /v1/readiness", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	mux.HandleFunc("POST /v1/users", cfg.handlerUsersCreate)
	mux.HandleFunc("GET /v1/users", cfg.authMiddleware(cfg.handlerUsersGet))

	mux.HandleFunc("POST /v1/feeds", cfg.authMiddleware(cfg.handlerFeedsCreate))
	mux.HandleFunc("GET /v1/feeds", cfg.handlerFeedsAllGet)

	mux.HandleFunc("POST /v1/feed_follows", cfg.authMiddleware(cfg.handlerFeedFollowCreate))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowId}", cfg.handlerFeedFollowDelete)
	mux.HandleFunc("GET /v1/feed_follows", cfg.authMiddleware(cfg.handlerFeedFollowGet))

	mux.HandleFunc("GET /v1/posts", cfg.authMiddleware(cfg.handlerPostsGet))

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
