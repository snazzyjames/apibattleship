package models

type Game struct {
	Id         string
	Players    map[string]Player
	PlayerTurn string
	Phase      string
}

type Games []*Game
