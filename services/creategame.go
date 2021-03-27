package services

import (
	"log"
	"math/rand"

	"github.com/snazzyjames/apibattleship/constants"
	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/services/util"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func CreateGame(p1 string, p2 string) *models.Game {
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

	return &game
}

func createBoard() models.Board {
	var boardSizeX, boardSizeY = util.GetBoardSize()

	Board := make([][]byte, boardSizeX)
	for row := range Board {
		Board[row] = make([]byte, boardSizeY)
	}
	log.Print(util.PrintBoard(Board))
	return Board
}

func createFleet() models.Ships {
	// We make a map of pointers to ship addresses so that we can check/update ships via pointer
	fleet := make(models.Ships)

	fleet["carrier"] = &models.Ship{
		Name:      "carrier",
		Mask:      constants.ShipCarrier,
		HitPoints: 5,
		Length:    5,
		Placed:    false,
	}
	fleet["battleship"] = &models.Ship{
		Name:      "battleship",
		Mask:      constants.ShipBattleship,
		HitPoints: 4,
		Length:    4,
		Placed:    false,
	}
	fleet["cruiser"] = &models.Ship{
		Name:      "cruiser",
		Mask:      constants.ShipCruiser,
		HitPoints: 3,
		Length:    3,
		Placed:    false,
	}
	fleet["submarine"] = &models.Ship{
		Name:      "submarine",
		Mask:      constants.ShipSubmarine,
		HitPoints: 3,
		Length:    3,
		Placed:    false,
	}
	fleet["destroyer"] = &models.Ship{
		Name:      "destroyer",
		Mask:      constants.ShipDestroyer,
		HitPoints: 2,
		Length:    2,
		Placed:    false,
	}

	return fleet
}
