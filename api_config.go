package main

import "sjohn/blog_aggregator/internal/database"

type apiConfig struct {
	DB *database.Queries
}
