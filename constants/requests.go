package constants

type NewGameRequest struct {
	PlayerOne string `json:"player_one"`
	PlayerTwo string `json:"player_two"`
}

type SetupGameRequest struct {
	Ship       string `json:"ship"`
	Coordinate string `json:"coordinate"`
	Direction  string `json:"direction"`
	Player     string `json:"player"`
}
