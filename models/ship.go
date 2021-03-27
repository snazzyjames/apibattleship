package models

type Ship struct {
	Name      string `json:"name"`
	Mask      byte   `json:"mask"`
	HitPoints int    `json:"hitpoints"`
	Length    int    `json:"length"`
	Placed    bool   `json:"placed"`
}

type Ships map[string]*Ship
