package main

import (
	"context"
	"net/http"
	"sjohn/blog_aggregator/internal/database"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (cfg *apiConfig) handlerPostUsers(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:"name"`
	}

	parameters := params{}
	if err := decodeRequestParams(r, &parameters); err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	ctx := context.Background()
	user, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		Name:      parameters.Name,
	})
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, user)
}
