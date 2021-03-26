package create

const boardSize = 10

func CreateBoard() [][]byte {
	Board := make([][]byte, boardSize)
	for row := range Board {
		Board[row] = make([]byte, boardSize)
	}
	return Board
}
