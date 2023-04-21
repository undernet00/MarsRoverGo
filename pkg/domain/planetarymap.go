package rover

type PlanetaryMap interface {
	IsValid(xcoord, ycoord int) bool
}

type Map struct {
	width   int
	height  int
	mapGrid [][]int
}

func (m *Map) IsValid(xcoord, ycoord int) bool {
	if xcoord < 0 || xcoord >= len(m.mapGrid) {
		return false
	}

	if ycoord < 0 || ycoord >= len(m.mapGrid[0]) {
		return false
	}

	return true
}

func NewPlanetaryMap(width, height int) *PlanetaryMap {
	//TODO: Validate parameters
	m := NewMap(width, height)
	newPm := PlanetaryMap(m)
	return &newPm
}

func NewMap(width, height int) *Map {
	newMap := Map{
		width:  width,
		height: height,
	}

	newMap.mapGrid = initializeGrid(width, height)

	return &newMap
}

func initializeGrid(width, height int) [][]int {
	newGrid := make([][]int, width)
	for x := 0; x < width; x++ {
		newGrid[x] = make([]int, height)
	}
	return newGrid
}
