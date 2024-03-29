package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panicln(err)
	}

	port := os.Getenv("PORT")
	runServer(port)
}
