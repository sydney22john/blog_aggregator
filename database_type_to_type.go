package main

import (
	"sjohn/blog_aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	UserID        uuid.UUID  `json:"user_id"`
	Url           string     `json:"url"`
	Name          string     `json:"name"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	var lastFetched *time.Time = nil
	if feed.LastFetchedAt.Valid {
		lastFetched = &feed.LastFetchedAt.Time
	}
	localFeed := Feed{
		ID:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		UserID:        feed.UserID,
		Url:           feed.Url,
		Name:          feed.Name,
		LastFetchedAt: lastFetched,
	}
	return localFeed
}
