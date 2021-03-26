package util

import "github.com/snazzyjames/apibattleship/models"

func Print(b models.Board) (str string) {
	var boardSizeX, boardSizeY = GetBoardSize()

	const (
		playerShot  = 1 << iota // 00000001
		playerHit               // 00000010
		opponentHit             // 00000100

		shipMask       = (1<<8 - 1) &^ (1<<5 - 1) // 11100000
		shipCarrier    = 1 << 5
		shipBattleship = 2 << 5
		shipDestroyer  = 3 << 5
		shipSubmarine  = 4 << 5
		shipPatrolBoat = 5 << 5
	)

	str += "\n  1 2 3 4 5 6 7 8 9 10\n"

	for y := 0; y < boardSizeY; y++ {
		str += string('A'+y) + " "
		for x := 0; x < boardSizeX; x++ {
			switch {
			case b[x][y]&playerHit > 0:
				str += "X "
			case b[x][y]&playerShot > 0:
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
