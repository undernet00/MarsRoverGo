package main

import (
	"fmt"
	planetarymap "github.com/undernet00/MarsRoverGo/pkg/domain"
)

func main() {
	p := planetarymap.NewPlanetaryMap(4, 3)

	fmt.Println(p)
}
