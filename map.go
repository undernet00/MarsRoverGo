package main

const (
	North string = "N"
	East         = "E"
	South        = "S"
	West         = "W"
)

type Map struct {
	width  int
	height int
}

func (m *Map) IsOut(xcoord, ycoord int) bool {

	return true
}

func NewMap(width, height int) *Map {
	newMap := Map{
		width:  width,
		height: height,
	}

	return &newMap
}
