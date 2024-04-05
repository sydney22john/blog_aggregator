package main

import (
	"context"
	"net/http"
	"sjohn/blog_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) authMiddleware(handler authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := getHeaderValue(r.Header.Get("Authorization"), "ApiKey ")
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		user, err := cfg.DB.SelectUserByApikey(context.Background(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "unauthenticated user")
			return
		}

		handler(w, r, user)
	})
}
