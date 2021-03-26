package models

type Game struct {
	Id         string   `json:"id"`
	Players    []Player `json:"players"`
	PlayerTurn int      `json:"playerturn"`
	Phase      string   `json:"phase"`
}
