package main

import (
	"context"
	"net/http"
	"sjohn/blog_aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	parameters := params{}
	if err := decodeRequestParams(r, &parameters); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        id,
		FeedID:    parameters.FeedId,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, feedFollow)
}
