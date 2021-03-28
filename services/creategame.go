package services

import (
	"log"

	"github.com/snazzyjames/apibattleship/constants"
	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/responses"
	"github.com/snazzyjames/apibattleship/util"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func CreateGame(p1 string, p2 string) (*models.Game, responses.NewGameResponse, error) {
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
	game.PlayerTurn = "p1"
	game.Phase = "setup"

	util.PrintStats(&game)

	return &game, responses.NewGameResponse{
		SessionId: game.Id,
		Phase:     game.Phase,
		Player:    game.Players[game.PlayerTurn].Name,
	}, nil
}

func createBoard() models.Board {
	var boardSizeX, boardSizeY = util.GetBoardSize()

	Board := make([][]byte, boardSizeX)
	for row := range Board {
		Board[row] = make([]byte, boardSizeY)
	}
	return Board
}

func createFleet() models.Ships {
	// We make a map of pointers (models.Ships)
	// to ship addresses so that we can check/update ships
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
