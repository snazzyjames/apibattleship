package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/snazzyjames/apibattleship/models"
)

var Games models.Games

func main() {
	router := Router()

	rand.Seed(time.Now().UnixNano())

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Printf("Listening on port %s", os.Getenv("APP_PORT"))

	// http.ListenAndServe opens the server port, and blocks forever waiting for clients.
	// If it fails to open the port, the log.Fatal call will report the problem and exit the program.
	log.Fatal(http.ListenAndServe(os.Getenv("APP_PORT"), router))
}
