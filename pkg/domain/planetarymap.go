package rover

type PlanetaryMap interface {
	IsValid(xcoord, ycoord int) bool
}

// IsValid validates a pair of x and y coordinates checking against map's width and height.
func (m *Map) IsValid(xCoordinate, yCoordinate int) bool {

	if xCoordinate < 0 || xCoordinate >= m.width {
		return false
	}

	if yCoordinate < 0 || yCoordinate >= m.height {
		return false
	}

	return true
}

func NewPlanetaryMap(width, height int) *PlanetaryMap {

	m := NewMap(width, height)
	newPm := PlanetaryMap(m)
	return &newPm
}

type Map struct {
	width  int
	height int
}

func NewMap(width, height int) *Map {

	newMap := Map{
		width:  width,
		height: height,
	}

	return &newMap
}
