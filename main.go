package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"sjohn/blog_aggregator/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const processFeeds = 3

func main() {
	randDBValues := flag.Bool("rand-db-values", false, "generate random values in the DB")
	if *randDBValues {
		log.Println("TODO later")
	}
	if err := godotenv.Load(); err != nil {
		log.Panicln(err)
	}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Panicln(err)
	}
	// checking for a valid connection
	if err = db.Ping(); err != nil {
		log.Panicln(err)
	}
	dbQueries := database.New(db)
	cfg := apiConfig{
		DB: dbQueries,
	}

	go cfg.blogAggregatorWorker(processFeeds)

	runServer(cfg, port)
}
