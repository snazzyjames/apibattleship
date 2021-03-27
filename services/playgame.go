package services

import (
	"fmt"
	"log"

	"github.com/snazzyjames/apibattleship/constants"
	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/util"
)

func PlayGame(game *models.Game, request constants.PlayGameRequest) (constants.PlayGameResponse, error) {
	if game.PlayerTurn != request.Player {
		log.Printf("error: Not %s's turn. Player turn is %s", request.Player, game.PlayerTurn)
		return constants.PlayGameResponse{
			Result:     "not_your_turn",
			NextPlayer: game.PlayerTurn,
		}, fmt.Errorf("error: Not %s's turn. Player turn is %s", request.Player, game.PlayerTurn)
	}

	if game.Phase != "setup" {
		log.Printf("cannot setup game, game phase is %v", game.Phase)
		return constants.PlayGameResponse{
			Result:     "",
			NextPlayer: game.PlayerTurn,
		}, fmt.Errorf("cannot setup game, game phase is %v", game.Phase)
	}

	players := game.Players
	var board models.Board
	var ships map[string]*models.Ship
	if (players)["p1"].Name == request.Player {
		board = (players)["p2"].Board
		ships = (players)["p2"].Ships
	} else {
		board = (players)["p1"].Board
		ships = (players)["p1"].Ships
	}

	x, y, err := util.ParsePosition(request.Coordinate)
	if err != nil {
		log.Println(err)
	}

	result := fire(&board, ships, x, y)

	// Print stats is used for debugging.  It will print the game players' ship statuses and boards
	util.PrintStats(game)

	allSunk := checkIfAllShipsSunk(ships)
	if allSunk {
		game.Phase = "game_over"
		result = "hit_good_game"
	}

	if game.PlayerTurn == game.Players["p1"].Name {
		game.PlayerTurn = game.Players["p2"].Name
	} else {
		game.PlayerTurn = game.Players["p1"].Name
	}

	return constants.PlayGameResponse{
		Result:     result,
		NextPlayer: game.PlayerTurn,
	}, nil
}

func checkIfAllShipsSunk(ships models.Ships) bool {
	var allSunk = true
	for _, ship := range ships {
		if ship.HitPoints > 0 {
			allSunk = false
		}
	}
	return allSunk
}

func fire(board *models.Board, ships models.Ships, x int, y int) string {
	shipPresent := (*board)[x][y] & constants.ShipMask
	var result = "miss"
	if shipPresent > 0 {
		// If we've already hit the ship at the coordinate, we'll just tell the user it's a hit but won't
		// decrement the ship's hitpoints.  We could check if the ship has been hit already and return an
		// "already_hit" message as a future improvement
		for _, playerShip := range ships {
			if (*board)[x][y] == playerShip.Mask && playerShip.HitPoints > 0 {
				playerShip.HitPoints -= 1
				(*board)[x][y] |= constants.Hit
			}

			// We could check if the ship has already been sunk and return an "already_sunk" message
			// as a future improvement, might need to add a "Sunk" flag on the ship model
			if playerShip.HitPoints == 0 {
				return "hit_sunk"
			}
		}
		return "hit"
	} else {
		(*board)[x][y] |= constants.Shot
	}
	return result
}
