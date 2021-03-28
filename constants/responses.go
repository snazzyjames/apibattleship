package constants

type NewGameResponse struct {
	SessionId string `json:"session_id"`
	Phase     string `json:"phase"`
	Player    string `json:"player"`
}

type GetSessionResponse struct {
	Phase   string    `json:"phase"`
	Players [2]string `json:"players"`
}

type SetupGameResponse struct {
	Placed     string `json:"placed"`
	NextPlayer string `json:"next_player,omitempty"`
	Phase      string `json:"phase"`
}

type PlayGameResponse struct {
	Result     string `json:"result"`
	NextPlayer string `json:"next_player"`
}
