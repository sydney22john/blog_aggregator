package main

import (
	"context"
	"net/http"
	"sjohn/blog_aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerPostUsers(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:"name"`
	}

	parameters := params{}
	if err := decodeRequestParams(r, &parameters); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userUuid, err := uuid.NewRandom()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ctx := context.Background()
	user, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        userUuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		Name:      parameters.Name,
		ApiKey:    createApiKey(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, user)
}
