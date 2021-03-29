package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/requests"
	"github.com/snazzyjames/apibattleship/responses"
	"github.com/snazzyjames/apibattleship/services"
)

/*
Handlers (i.e. Controllers) to check and marshall incoming requests, set content type, call service functions,
and write responses
*/
func NewGame(w http.ResponseWriter, r *http.Request) {
	w = setContentType(w)
	var request requests.NewGameRequest
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
	// TODO: validate and sanitize input params using Golang best practices
	newGame, response, err := services.CreateGame(request.PlayerOne, request.PlayerTwo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Games = append(Games, newGame)
	json.NewEncoder(w).Encode(response)
}

func SetupGame(w http.ResponseWriter, r *http.Request) {
	w = setContentType(w)
	var request requests.SetupGameRequest
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
	// TODO: validate and sanitize input params using Golang best practices
	vars := mux.Vars(r)
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

func PlayGame(w http.ResponseWriter, r *http.Request) {
	w = setContentType(w)
	var request requests.PlayGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if request.Player == "" || request.Coordinate == "" {
		http.Error(w, "unrecognized json", http.StatusBadRequest)
		return
	}
	// TODO: validate and sanitize input params using Golang best practices
	vars := mux.Vars(r)
	gameId := vars["sessionId"]
	game := getGameById(gameId)

	if game == nil {
		w.WriteHeader(404)
		return
	}
	response, err := services.PlayGame(game, request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(response)
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	w = setContentType(w)

	// TODO: future improvement: validate and sanitize input params using Golang best practices
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

	json.NewEncoder(w).Encode(responses.GetGameResponse{
		Phase:   game.Phase,
		Players: players,
	})
}

func setContentType(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	return w
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
