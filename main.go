package main

import (
	"log"
	"net/http"

	"github.com/snazzyjames/apibattleship/models"
)

var Sessions []models.Game

func main() {
	router := BattleshipRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
