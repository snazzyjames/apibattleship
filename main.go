package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/joho/godotenv"

	"github.com/snazzyjames/apibattleship/models"
)

var Sessions []models.Game

func main() {
	router := Router()

	rand.Seed(time.Now().UnixNano())

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// http.ListenAndServe opens the server port, and blocks forever waiting for clients.
	// If it fails to open the port, the log.Fatal call will report the problem and exit the program.
	log.Fatal(http.ListenAndServe(":8080", router))
}
