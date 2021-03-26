package services

import (
	"log"
	"math/rand"

	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/services/util"

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
		Ships: createFleet(),
		Board: createBoard(),
	}
	player2 := models.Player{
		Id:    2,
		Name:  p2,
		Ships: createFleet(),
		Board: createBoard(),
	}
	players := make(map[string]models.Player)
	players["p1"] = player1
	players["p2"] = player2
	game.Players = players

	coinFlip := rand.Intn(2) == 0
	if coinFlip == true {
		game.PlayerTurn = player1.Name
	} else {
		game.PlayerTurn = player2.Name
	}

	game.Phase = "setup"

	return game
}

func createBoard() models.Board {
	var boardSizeX, boardSizeY = util.GetBoardSize()

	Board := make([][]byte, boardSizeX)
	for row := range Board {
		Board[row] = make([]byte, boardSizeY)
	}
	log.Print(util.Print(Board))
	return Board
}

func createFleet() map[string]models.Ship {
	fleet := make(map[string]models.Ship)

	fleet["carrier"] = models.Ship{
		Name:      "carrier",
		Mask:      1 << 5,
		HitPoints: 5,
	}
	fleet["battleship"] = models.Ship{
		Name:      "battleship",
		Mask:      2 << 5,
		HitPoints: 4,
	}
	fleet["cruiser"] = models.Ship{
		Name:      "cruiser",
		Mask:      3 << 5,
		HitPoints: 3,
	}
	fleet["submarine"] = models.Ship{
		Name:      "submarine",
		Mask:      4 << 5,
		HitPoints: 3,
	}
	fleet["destroyer"] = models.Ship{
		Name:      "destroyer",
		Mask:      5 << 5,
		HitPoints: 2,
	}

	return fleet
}
