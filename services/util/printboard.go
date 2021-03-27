package util

import (
	"github.com/snazzyjames/apibattleship/constants"
	"github.com/snazzyjames/apibattleship/models"
)

func PrintBoard(b models.Board) (str string) {
	var boardSizeX, boardSizeY = GetBoardSize()

	str += "\n  0 1 2 3 4 5 6 7 8 9\n"

	for y := 0; y < boardSizeY; y++ {
		str += string('A'+y) + " "
		for x := 0; x < boardSizeX; x++ {
			switch {
			case b[x][y] == constants.ShipCarrier:
				str += "C "
			case b[x][y] == constants.ShipBattleship:
				str += "B "
			case b[x][y] == constants.ShipCruiser:
				str += "K "
			case b[x][y] == constants.ShipSubmarine:
				str += "S "
			case b[x][y] == constants.ShipDestroyer:
				str += "D "
			case b[x][y]&constants.Hit > 0:
				str += "X "
			case b[x][y]&constants.Shot > 0:
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
