package models

type Ship struct {
	Name      string `json:"name"`
	Mask      byte   `json:"mask"`
	HitPoints int    `json:"hitpoints"`
}
