package main

import (
	"fmt"
	planetarymap "github.com/undernet00/MarsRoverGo/pkg/domain"
)

func main() {
	p := planetarymap.NewPlanetaryMap(4, 4)
	r := planetarymap.NewRover(*p)

	output, err := r.Travel(0, 3, planetarymap.South, "AAALAAALAAA")

	fmt.Println(output)
	fmt.Println(err)

}
