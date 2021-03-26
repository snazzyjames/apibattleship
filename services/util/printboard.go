package util

import (
	"github.com/snazzyjames/apibattleship/constants"
	"github.com/snazzyjames/apibattleship/models"
)

func Print(b models.Board) (str string) {
	var boardSizeX, boardSizeY = GetBoardSize()

	str += "\n  1 2 3 4 5 6 7 8 9 10\n"

	for y := 0; y < boardSizeY; y++ {
		str += string('A'+y) + " "
		for x := 0; x < boardSizeX; x++ {
			switch {
			case b[x][y]&constants.PlayerHit > 0:
				str += "X "
			case b[x][y]&constants.PlayerShot > 0:
				str += "O "
			default:
				str += "- "
			}
		}
		str += "\n"
	}
	str += "\n"
	return
}
