package models

type Ship struct {
	Name      string
	Mask      byte
	HitPoints int
	Length    int
	Placed    bool
}

type Ships map[string]*Ship
