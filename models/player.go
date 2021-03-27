package models

type Player struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Ships Ships  `json:"ships"`
	Board Board  `json:"board"`
}
