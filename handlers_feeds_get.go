package main

import (
	"context"
	"net/http"
)

func (cfg *apiConfig) handlerFeedsAllGet(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, feeds)
}
