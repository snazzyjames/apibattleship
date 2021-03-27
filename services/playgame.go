package services

import (
	"errors"
	"log"

	"github.com/snazzyjames/apibattleship/constants"
	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/services/util"
)

func PlayGame(game *models.Game, coordinate string, player string) (result string, nextPlayer string, err error) {
	if game.PlayerTurn != player {
		log.Printf("error: Not %s's turn. Player turn is %s", player, game.PlayerTurn)
		return "not_your_turn", game.PlayerTurn, errors.New("Invalid player turn")
	}

	players := game.Players
	var board models.Board
	var ships map[string]*models.Ship
	if (players)["p1"].Name == player {
		board = (players)["p2"].Board
		ships = (players)["p2"].Ships
	} else {
		board = (players)["p1"].Board
		ships = (players)["p1"].Ships
	}

	x, y, err := util.ParsePosition(coordinate)
	if err != nil {
		log.Println(err)
	}

	result = fire(&board, ships, x, y)

	// Debug logic for outputting result of shots
	log.Println("Player 1 Board/Ship Status")
	for _, p1ship := range players["p1"].Ships {
		log.Printf("Name: %v HP: %v/%v", p1ship.Name, p1ship.HitPoints, p1ship.Length)
	}
	log.Println(util.PrintBoard(players["p1"].Board))

	log.Println("Player 2 Board/Ship Status")
	for _, p2ship := range players["p2"].Ships {
		log.Printf("Name: %v HP: %v/%v", p2ship.Name, p2ship.HitPoints, p2ship.Length)
	}
	log.Println(util.PrintBoard(players["p2"].Board))

	// End Debug logic

	allSunk := checkIfAllShipsSunk(ships)
	if allSunk {
		result = "hit_good_game"
	}

	return result, "james", nil
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
