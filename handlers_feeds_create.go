package main

import (
	"context"
	"net/http"
	"sjohn/blog_aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		URL  string `json:"url"`
		Name string `json:"name"`
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

	feed, err := cfg.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		UserID:    user.ID,
		Url:       parameters.URL,
		Name:      parameters.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err = uuid.NewRandom()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	feedFollow, err := cfg.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        id,
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	type resp struct {
		Feed       Feed               `json:"feed"`
		FeedFollow database.UsersFeed `json:"feed_follow"`
	}
	respondWithJson(w, http.StatusOK, resp{
		Feed:       databaseFeedToFeed(feed),
		FeedFollow: feedFollow,
	})
}
