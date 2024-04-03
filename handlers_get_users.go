package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetUsersByApikey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := getHeaderValue(r.Header.Get("Authorization"), "ApiKey ")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := context.Background()
	user, err := cfg.DB.SelectUserByApikey(ctx, apiKey)

	if user.ID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "user not found")
		return
	}

	respondWithJson(w, http.StatusOK, user)
}
