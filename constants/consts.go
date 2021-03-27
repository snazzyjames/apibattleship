package constants

const (
	PlayerShot  = 1 << iota // 00000001
	PlayerHit               // 00000010
	OpponentHit             // 00000100

	// The first 3 bits represent the type of ship, so the ship mask
	// is what we can use to identify bytes that are a ship
	// TODO: More description about what this cryptic bit operation does
	ShipMask       = (1<<8 - 1) &^ (1<<5 - 1) // 11100000
	ShipCarrier    = 1 << 5
	ShipBattleship = 2 << 5
	ShipCruiser    = 3 << 5
	ShipSubmarine  = 4 << 5
	ShipDestroyer  = 5 << 5
	Empty          = 0 << 5
)
