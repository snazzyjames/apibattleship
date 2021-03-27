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
	w = setContentType(w)
	var request constants.NewGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Panic(err.Error())
		return
	}
	if request.PlayerOne == "" || request.PlayerTwo == "" {
		message := "Two players required"
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	newGame, response, err := services.CreateGame(request.PlayerOne, request.PlayerTwo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Games = append(Games, newGame)
	json.NewEncoder(w).Encode(response)
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
	if request.Player == "" || request.Coordinate == "" || request.Direction == "" || request.Ship == "" {
		http.Error(w, "unrecognized json", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	// TODO: Validate and sanitize params
	gameId := vars["sessionId"]
	game := getGameById(gameId)
	if game == nil {
		w.WriteHeader(404)
		return
	}
	response, err := services.SetupGame(game, request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(response)
	return
}

func PlaySession(w http.ResponseWriter, r *http.Request) {
	w = setContentType(w)
	var request constants.PlayGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if request.Player == "" || request.Coordinate == "" {
		http.Error(w, "unrecognized json", http.StatusBadRequest)
		return
	}
	// TODO: Validate and sanitize params
	vars := mux.Vars(r)
	gameId := vars["sessionId"]
	game := getGameById(gameId)

	if game == nil {
		w.WriteHeader(404)
		return
	}
	if game.Phase == "setup" || game.Phase == "game_over" {
		http.Error(w, "game is not in play phase", http.StatusBadRequest)
		return
	}
	response, err := services.PlayGame(game, request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(response)
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	w = setContentType(w)
	// TODO: Validate and sanitize params
	vars := mux.Vars(r)
	gameId := vars["sessionId"]

	game := getGameById(gameId)
	if game == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var players = [2]string{
		game.Players["p1"].Name,
		game.Players["p2"].Name,
	}

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
