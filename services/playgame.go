package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/snazzyjames/apibattleship/constants"
	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/util"
)

func PlayGame(game *models.Game, request constants.PlayGameRequest) (constants.PlayGameResponse, error) {
	if game.Phase != "play" {
		result := "not_your_turn"
		if game.Phase == "game_over" {
			result = "game_over"
		}
		log.Printf("cannot play game, game phase is %v", game.Phase)
		return constants.PlayGameResponse{
			Result:     result,
			NextPlayer: game.Players[game.PlayerTurn].Name,
		}, fmt.Errorf("cannot play game, game phase is %v", game.Phase)
	}

	if game.Players[game.PlayerTurn].Name != request.Player {
		log.Printf("error: Not %s's turn. Player turn is %s", request.Player, game.PlayerTurn)
		return constants.PlayGameResponse{
			Result:     "not_your_turn",
			NextPlayer: game.PlayerTurn,
		}, fmt.Errorf("error: Not %s's turn. Player turn is %s", request.Player, game.PlayerTurn)
	}

	players := game.Players
	board := (players)[game.PlayerTurn].Board
	ships := (players)[game.PlayerTurn].Ships

	x, y, err := util.ParsePosition(request.Coordinate)
	if err != nil {
		// If the position is invalid we return a result of miss with an error code of 500,
		// may want to add to the contract for when errors like this happen
		log.Println(err)
		return constants.PlayGameResponse{
			Result:     "miss",
			NextPlayer: game.Players[game.PlayerTurn].Name,
		}, errors.New("failed parsing position")
	}

	result := fire(&board, ships, x, y)

	// Print stats is used for debugging.  It will print the game players' ship statuses and boards
	util.PrintStats(game)

	if game.PlayerTurn == "p1" {
		game.PlayerTurn = "p2"
	} else {
		game.PlayerTurn = "p1"
	}

	allSunk := checkIfAllShipsSunk(ships)
	if allSunk {
		game.Phase = "game_over"
		result = "hit_good_game"
	}

	return constants.PlayGameResponse{
		Result:     result,
		NextPlayer: game.Players[game.PlayerTurn].Name,
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
	if shipPresent > 0 {
		// If a player has already hit the ship at the coordinate, just tell the user it's a hit but don't
		// decrement the ship's hitpoints.  The code could check if the ship has been hit already and return an
		// "already_hit" message as a future improvement
		for _, playerShip := range ships {
			if (*board)[x][y] == playerShip.Mask && playerShip.HitPoints > 0 {
				playerShip.HitPoints -= 1
				(*board)[x][y] |= constants.Hit
				// We could check if the ship has already been sunk and return an "already_sunk" message
				// as a future improvement, might need to add a "Sunk" flag on the ship model
				if playerShip.HitPoints == 0 {
					return "hit_sunk"
				}
				break
			}
		}
		return "hit"
	} else {
		(*board)[x][y] |= constants.Shot
	}
	return "miss"
}
