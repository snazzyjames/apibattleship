package util

// IsValid returns true if the given coordinates are a location on the board, else false
func IsValidCoord(x, y int) bool {
	boardSizeX, boardSizeY := GetBoardSize()
	if x < 0 || x >= boardSizeX {
		return false
	}
	if y < 0 || y >= boardSizeY {
		return false
	}
	return true
}
