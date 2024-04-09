package main

import (
	"context"
	"log"
	"math/rand"
	"sjohn/blog_aggregator/internal/database"

	"github.com/google/uuid"
)

func randomDBValues(cfg *apiConfig) {
	userIds := make([]uuid.UUID, 0)
	// creates 3 to 5 random users
	for range rand.Intn(3) + 3 {
		userId, err := uuid.NewRandom()
		if err != nil {
			log.Panicln(err)
		}
		cfg.DB.CreateUser(context.Background(), database.CreateUserParams{
			ID: userId,
		})
		userIds = append(userIds, userId)
	}
}
