package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/snazzyjames/apibattleship/constants"
	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/services"
)

func setContentType(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	return w
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	// TODO: Add validation rules here to sanitize JSON request

	var request constants.NewGameRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Panic(err.Error())
		return
	} else if request.PlayerOne == "" || request.PlayerTwo == "" {
		message := "Two players required"
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	w = setContentType(w)
	newGame := services.CreateGame(request.PlayerOne, request.PlayerTwo)
	Games = append(Games, newGame)
	json.NewEncoder(w).Encode(Games)
}

func SetupSession(w http.ResponseWriter, r *http.Request) {
	w = setContentType(w)
	var request constants.SetupGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Panic(err.Error())
		return
	}
	vars := mux.Vars(r)
	gameId := vars["sessionId"]
	game := getGameById(gameId)
	response, err := services.SetupGame(game, request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(response)
	return
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	// TODO: Validate and sanitize params
	vars := mux.Vars(r)
	gameId := vars["sessionId"]

	game := getGameById(gameId)
	if game.Id == "" {
		w.WriteHeader(404)
		return
	}

	var players = [2]string{
		game.Players["p1"].Name,
		game.Players["p2"].Name,
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(constants.GetSessionResponse{
		Phase:   game.Phase,
		Players: players,
	})
}

func getGameById(gameId string) *models.Game {
	if len(Games) != 0 {
		for _, game := range Games {
			if game.Id == gameId {
				return game
			}
		}
	}
	return nil
}
