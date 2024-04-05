package main

import (
	"net/http"
	"sjohn/blog_aggregator/internal/database"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, http.StatusOK, user)
}
