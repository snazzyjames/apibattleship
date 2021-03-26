package models

type Player struct {
	Id    int             `json:"id"`
	Name  string          `json:"name"`
	Ships map[string]Ship `json:"ships"`
	Board [][]byte        `json:"board"`
}
