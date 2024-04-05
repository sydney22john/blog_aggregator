package main

import (
	"context"
	"net/http"
	"sjohn/blog_aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedFollowGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := cfg.DB.GetFeedFollows(context.Background(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, feeds)
}
