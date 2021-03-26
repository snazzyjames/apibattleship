package create

import (
	"log"

	"github.com/snazzyjames/apibattleship/models"
	"github.com/snazzyjames/apibattleship/services/util"
)

func CreateBoard() models.Board {
	var boardSizeX, boardSizeY = util.GetBoardSize()

	Board := make([][]byte, boardSizeX)
	for row := range Board {
		Board[row] = make([]byte, boardSizeY)
	}
	log.Print(util.Print(Board))
	return Board
}
