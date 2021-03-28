package util

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParsePosition(location string) (x int, y int, err error) {
	if location == "" {
		return 0, 0, errors.New("no location specified")
	}

	location = strings.TrimSpace(strings.ToLower(location))
	y = int(location[0] - 'a')          // first char of string in asii minus the ascii code for 'a'
	x, err = strconv.Atoi(location[1:]) // string to int conversion for remaining string bytes.
	if err != nil {
		return 0, 0, fmt.Errorf("invalid location %v", location)
	}
	if !IsValidCoord(x, y) {
		return 0, 0, fmt.Errorf("invalid location %v", location)
	}
	return
}
