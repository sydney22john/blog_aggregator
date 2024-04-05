package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request) {
	feedFollowIdString := r.PathValue("feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = cfg.DB.DeleteFeedFollow(context.Background(), feedFollowId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, nil)
}
