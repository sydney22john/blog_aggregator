package main

import (
	"database/sql"
	"log"
	"os"
	"sjohn/blog_aggregator/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panicln(err)
	}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Panicln(err)
	}

	if err = db.Ping(); err != nil {
		log.Panicln(err)
	}
	dbQueries := database.New(db)
	cfg := apiConfig{
		DB: dbQueries,
	}

	runServer(cfg, port)
}
