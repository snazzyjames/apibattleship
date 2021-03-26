package services

import (
	"log"

	"github.com/snazzyjames/apibattleship/constants"
	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/services/util"
)

func SetupGame(game models.Game, request constants.SetupGameRequest) constants.SetupGameResponse {
	if game.PlayerTurn != request.Player {
		return constants.SetupGameResponse{
			Placed:     "false",
			NextPlayer: game.PlayerTurn,
			Phase:      game.Phase,
		}
	}
	log.Println(util.Print(game.Players["p1"].Board))
	return constants.SetupGameResponse{}
}
