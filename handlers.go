package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/snazzyjames/apibattleship/services/create"
)

func setContentType(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	return w
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	type NewGameRequest struct {
		PlayerOne string `json:"player_one"`
		PlayerTwo string `json:"player_two"`
	}
	var request NewGameRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if request.PlayerOne == "" || request.PlayerTwo == "" {
		message := "Two players required"
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().UnixNano())
	w = setContentType(w)
	newGame := create.CreateGame(request.PlayerOne, request.PlayerTwo)
	Sessions = append(Sessions, newGame)
	json.NewEncoder(w).Encode(Sessions)
}
