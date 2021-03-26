package create

import "github.com/snazzyjames/apibattleship/models"

func CreateFleet() map[string]models.Ship {
	fleet := make(map[string]models.Ship)

	fleet["carrier"] = models.Ship{
		Name:      "carrier",
		Mask:      1 << 5,
		HitPoints: 5,
	}
	fleet["battleship"] = models.Ship{
		Name:      "battleship",
		Mask:      2 << 5,
		HitPoints: 4,
	}
	fleet["cruiser"] = models.Ship{
		Name:      "cruiser",
		Mask:      3 << 5,
		HitPoints: 3,
	}
	fleet["submarine"] = models.Ship{
		Name:      "submarine",
		Mask:      4 << 5,
		HitPoints: 3,
	}
	fleet["destroyer"] = models.Ship{
		Name:      "destroyer",
		Mask:      5 << 5,
		HitPoints: 2,
	}

	return fleet
}
