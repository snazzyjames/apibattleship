package services

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/snazzyjames/apibattleship/constants"
	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/services/util"
)

func SetupGame(game *models.Game, request constants.SetupGameRequest) constants.SetupGameResponse {
	if game.PlayerTurn != request.Player {
		log.Printf("error: Not %s's turn. Player turn is %s", request.Player, game.PlayerTurn)
		return constants.SetupGameResponse{
			Placed:     "false",
			NextPlayer: game.PlayerTurn,
			Phase:      game.Phase,
		}
	}

	players := game.Players

	var board models.Board
	var ships map[string]*models.Ship
	if (players)["p1"].Name == request.Player {
		board = (players)["p1"].Board
		ships = (players)["p1"].Ships
	} else {
		board = (players)["p2"].Board
		ships = (players)["p2"].Ships
	}

	x, y, err := util.ParsePosition(request.Coordinate)
	if err != nil {
		log.Println(err)
	}

	var ship *models.Ship
	for key, playerShip := range ships {
		if key == request.Ship {
			ship = playerShip
			break
		}
	}

	placed, err := placeShip(&board, x, y, ship, request.Direction)
	if err != nil {
		log.Println(err)
	}
	if placed {
		if game.PlayerTurn == game.Players["p1"].Name {
			game.PlayerTurn = game.Players["p2"].Name
		} else {
			game.PlayerTurn = game.Players["p1"].Name
		}
	}

	allPlaced := checkIfAllShipsPlaced(players["p1"].Ships, players["p2"].Ships)
	if allPlaced {
		game.Phase = "play"
	}

	log.Println(util.PrintBoard(players["p1"].Board))
	log.Println(util.PrintBoard(players["p2"].Board))

	return constants.SetupGameResponse{
		Placed:     strconv.FormatBool(placed),
		NextPlayer: game.PlayerTurn,
		Phase:      game.Phase,
	}
}

func checkIfAllShipsPlaced(ships1 models.Ships, ships2 models.Ships) bool {
	var allPlaced = true
	for _, ship := range ships1 {
		if ship.Placed == false {
			allPlaced = false
		}
	}
	for _, ship := range ships2 {
		if ship.Placed == false {
			allPlaced = false
		}
	}
	return allPlaced
}

func placeShip(board *models.Board, x int, y int, ship *models.Ship, direction string) (bool, error) {
	if ship.Placed == true {
		return false, errors.New("error: ship already placed")
	}
	var mx, my int
	switch direction {
	case "right":
		mx = 1
	case "down":
		my = 1
	}

	ix, iy := x, y // create copies of x and y to use for investigation of board
	for i := 0; i < ship.Length; i++ {
		if !util.IsValidCoord(ix, iy) {
			return false, errors.New("error: ship is off the board")
		}
		if (*board)[ix][iy]&constants.ShipMask > 0 {
			return false, fmt.Errorf("error: there is already a ship at %v", util.FormatPosition(ix, iy))
		}
		ix += mx
		iy += my
	}

	// Everything is good to go, place the ship on the board.
	for i := 0; i < ship.Length; i++ {
		(*board)[x][y] = ship.Mask
		x += mx
		y += my
	}

	(*ship).Placed = true
	return true, nil
}
