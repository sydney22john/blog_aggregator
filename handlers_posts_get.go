package main

import (
	"context"
	"net/http"
	"sjohn/blog_aggregator/internal/database"
	"strconv"
)

func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	limitPostsTo := 10
	if limit := r.URL.Query().Get("limit"); limit != "" {
		var err error
		limitPostsTo, err = strconv.Atoi(limit)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	posts, err := cfg.DB.GetPostsByUser(context.Background(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limitPostsTo),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, mapArray(posts, databasePostToPost))
}
