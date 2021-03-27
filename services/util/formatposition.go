package util

import "strconv"

// Takes x and y coordinates and formats them for the game board e.g. x=2,y=6 -> "b6"
// The cast of 'a'+y to a rune is done so that we get the ascii/unicode code point for that letter of the alphabet.
// We can then cast the resulting codepoint to a string to get the correct row letter
func FormatPosition(x, y int) string {
	return string(rune('a'+y)) + strconv.Itoa(x)
}
