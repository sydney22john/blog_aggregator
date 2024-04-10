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

type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	var description *string = nil
	if post.Description.Valid {
		description = &post.Description.String
	}
	var publishedAt *time.Time = nil
	if post.PublishedAt.Valid {
		publishedAt = &post.PublishedAt.Time
	}

	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: description,
		PublishedAt: publishedAt,
		FeedID:      post.FeedID,
	}
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
