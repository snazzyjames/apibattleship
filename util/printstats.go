package util

import (
	"log"

	"github.com/snazzyjames/apibattleship/models"
)

func PrintStats(game *models.Game) {
	players := game.Players
	log.Printf("Game: %v GamePhase: %v PlayerTurn: %v", game.Id, game.Phase, game.PlayerTurn)
	log.Println("---------------------")
	log.Printf("Player 1 (%v) Board/Ship Status", players["p1"].Name)
	for _, p1ship := range players["p1"].Ships {
		log.Printf("Name: %v HP: %v/%v Placed: %v", p1ship.Name, p1ship.HitPoints, p1ship.Length, p1ship.Placed)
	}
	log.Println(PrintBoard(players["p1"].Board))

	log.Printf("Player 2 (%v) Board/Ship Status", players["p2"].Name)
	for _, p2ship := range players["p2"].Ships {
		log.Printf("Name: %v HP: %v/%v Placed: %v", p2ship.Name, p2ship.HitPoints, p2ship.Length, p2ship.Placed)
	}
	log.Println(PrintBoard(players["p2"].Board))
}
