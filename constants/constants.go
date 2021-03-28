package constants

const (
	// The iota keyword represents successive integer constants 0, 1, 2,…
	Shot = 1 << iota // 00000001
	Hit              // 00000010

	// The first 3 bits represent the type of ship, so the ship mask
	// is what we can use to identify bytes that are a ship, by using & and seeing if the value is greater than one.
	ShipMask       = 1<<7 + 1<<6 + 1<<5 // 11100000
	ShipCarrier    = 1 << 5
	ShipBattleship = 2 << 5
	ShipCruiser    = 3 << 5
	ShipSubmarine  = 4 << 5
	ShipDestroyer  = 5 << 5
)
