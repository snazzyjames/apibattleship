package models

type Ship struct {
	Name      string       `json:"name"`
	Length    int          `json:"length"`
	Spaces    []coordinate `json:"spaces"`
	HitPoints int          `json:"hitpoints"`
}
