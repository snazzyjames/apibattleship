package models

type Player struct {
	Id    int
	Name  string
	Ships Ships
	Board Board
}
