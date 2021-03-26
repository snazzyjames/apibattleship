package util

import (
	"log"
	"os"
	"strconv"
)

func GetBoardSize() (int, int) {
	var x = os.Getenv("BOARD_SIZE_X")
	var y = os.Getenv("BOARD_SIZE_Y")

	if x == "" || y == "" {
		log.Panicf("No environment variables set for Board Size X or Y")
	}

	boardSizeX, err := strconv.Atoi(x)
	if err != nil {
		log.Panic(err)
	}

	boardSizeY, err := strconv.Atoi(x)
	if err != nil {
		log.Panic(err)
	}

	return boardSizeX, boardSizeY
}
