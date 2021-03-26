package models

type Game struct {
	Id         string            `json:"id"`
	Players    map[string]Player `json:"players"`
	PlayerTurn string            `json:"playerturn"`
	Phase      string            `json:"phase"`
}
