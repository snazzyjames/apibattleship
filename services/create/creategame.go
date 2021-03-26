package create

import (
	"log"
	"math/rand"

	"github.com/snazzyjames/apibattleship/models"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func CreateGame(p1 string, p2 string) models.Game {
	var game models.Game

	id, err := gonanoid.New(5)
	if err != nil {
		log.Fatal(err)
	}

	game.Id = id
	player1 := models.Player{
		Id:    1,
		Name:  p1,
		Ships: CreateFleet(),
		Board: CreateBoard(),
	}
	player2 := models.Player{
		Id:    2,
		Name:  p2,
		Ships: CreateFleet(),
		Board: CreateBoard(),
	}
	players := make(map[string]models.Player)
	players["p1"] = player1
	players["p2"] = player2
	game.Players = players

	coinFlip := rand.Intn(2) == 0
	if coinFlip == true {
		game.PlayerTurn = "p1"
	} else {
		game.PlayerTurn = "p2"
	}

	game.Phase = "setup"

	return game
}
