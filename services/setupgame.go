package services

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/snazzyjames/apibattleship/constants"
	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/util"
)

func SetupGame(game *models.Game, request constants.SetupGameRequest) (constants.SetupGameResponse, error) {
	if game.Phase != "setup" {
		log.Printf("cannot setup game, game phase is %v", game.Phase)
		return constants.SetupGameResponse{
			Placed:     "false",
			NextPlayer: game.Players[game.PlayerTurn].Name,
			Phase:      game.Phase,
		}, fmt.Errorf("cannot setup game, game phase is %v", game.Phase)
	}

	if game.Players[game.PlayerTurn].Name != request.Player {
		log.Printf("error: Not %s's turn. Player turn is %s", request.Player, game.Players[game.PlayerTurn].Name)
		return constants.SetupGameResponse{
			Placed:     "false",
			NextPlayer: game.Players[game.PlayerTurn].Name,
			Phase:      game.Phase,
		}, errors.New("incorrect player turn")
	}

	players := game.Players
	board := (players)[game.PlayerTurn].Board
	ships := (players)[game.PlayerTurn].Ships

	x, y, err := util.ParsePosition(request.Coordinate)
	if err != nil {
		log.Println(err)
		return constants.SetupGameResponse{
			Placed:     "false",
			NextPlayer: game.Players[game.PlayerTurn].Name,
			Phase:      game.Phase,
		}, errors.New("failed parsing position")
	}

	ship := ships[request.Ship]
	if ship == nil {
		log.Println("invalid ship")
		return constants.SetupGameResponse{
			Placed:     "false",
			NextPlayer: game.Players[game.PlayerTurn].Name,
			Phase:      game.Phase,
		}, errors.New("invalid ship")
	}

	placed, err := placeShip(&board, x, y, ship, request.Direction)
	if err != nil {
		log.Println(err)
		return constants.SetupGameResponse{
			Placed:     "false",
			NextPlayer: game.Players[game.PlayerTurn].Name,
			Phase:      game.Phase,
		}, err
	}

	if placed {
		if game.PlayerTurn == "p1" {
			game.PlayerTurn = "p2"
		} else {
			game.PlayerTurn = "p1"
		}
	}

	util.PrintStats(game)

	allPlaced := checkIfAllShipsPlaced(players["p1"].Ships, players["p2"].Ships)
	if allPlaced {
		game.Phase = "play"
		return constants.SetupGameResponse{
			Placed: strconv.FormatBool(placed),
			Phase:  game.Phase,
		}, nil
	}

	return constants.SetupGameResponse{
		Placed:     strconv.FormatBool(placed),
		NextPlayer: game.Players[game.PlayerTurn].Name,
		Phase:      game.Phase,
	}, nil
}

func checkIfAllShipsPlaced(p1ships models.Ships, p2ships models.Ships) bool {
	var allPlaced = true
	for _, ship := range p1ships {
		if ship.Placed == false {
			allPlaced = false
		}
	}
	for _, ship := range p2ships {
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

	if direction != "right" && direction != "down" {
		return false, errors.New("error: invalid direction")
	}
	switch direction {
	case "right":
		mx = 1
	case "down":
		my = 1
	}

	ix, iy := x, y // create copies of x and y for investigating board
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

	// Place ship on board after checking it won't be placed out of bounds or on top of another ship
	for i := 0; i < ship.Length; i++ {
		(*board)[x][y] = ship.Mask
		x += mx
		y += my
	}

	(*ship).Placed = true
	return true, nil
}
