package rover

type CardinalPoint string

const (
	North CardinalPoint = "N"
	East  CardinalPoint = "E"
	South CardinalPoint = "S"
	West  CardinalPoint = "W"
)

type Command string

const (
	Advance Command = "A"
	Left    Command = "L"
	Right   Command = "R"
)

type Rover struct {
	listOfCommands     string
	currentOrientation CardinalPoint
	currentX           int
	currentY           int
	navigationMap      PlanetaryMap
}

func NewRover(navigationMap PlanetaryMap, initialX int, initialY int, initialOrientation CardinalPoint, listOfCommands string) *Rover {
	//TODO: Validate all parameters

	newRover := Rover{
		navigationMap:      navigationMap,
		currentX:           initialX,
		currentY:           initialY,
		listOfCommands:     listOfCommands,
		currentOrientation: initialOrientation,
	}

	return &newRover
}

func convertStringToCommands(listOfCommands string) ([]Command, error) {

	return nil, nil
}
