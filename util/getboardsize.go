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
		log.Fatalf("No environment variables set for BOARD_SIZE_X or BOARD_SIZE_Y, please check .env")
	}

	boardSizeX, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}

	boardSizeY, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
	}

	return boardSizeX, boardSizeY
}
