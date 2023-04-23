package rover

import (
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

type CardinalPoint string

const (
	North CardinalPoint = "N"
	East  CardinalPoint = "E"
	South CardinalPoint = "S"
	West  CardinalPoint = "W"
)

// IsValid validates that the value of the CardinalPoint is one of the four possible values.
func (cp CardinalPoint) IsValid() bool {
	return cp == North || cp == East || cp == South || cp == West
}

type Command string

const (
	Advance Command = "A"
	Left    Command = "L"
	Right   Command = "R"
)

// IsValid validates that the value of the Command is one of the tree possible values.
//
//	-A for Advance.
//	-L for Left.
//	-R for Right.
func (c Command) IsValid() bool {

	return c == Advance || c == Left || c == Right
}

type Rover struct {
	currentOrientation CardinalPoint
	currentX           int
	currentY           int
	navigationMap      PlanetaryMap
}

func NewRover(navigationMap PlanetaryMap) *Rover {

	if navigationMap == nil {
		return nil
	}

	newRover := Rover{
		navigationMap: navigationMap,
	}

	return &newRover
}

// Travel will take an initial x and y position, an initial orientation and a list of commands.
// Then will try to simulate the rover's travel on the map and return a formatted string with the result.
func (r *Rover) Travel(initialX int, initialY int, initialOrientation CardinalPoint, listOfCommands string) (string, error) {

	if r == nil {
		return "", errors.New("Rover was not initialized\n")
	}

	commands, err := convertStringToCommands(listOfCommands)
	if err != nil {
		return "", err
	}

	if !r.navigationMap.IsValid(initialX, initialY) {
		return "", errors.New(fmt.Sprintf("(%v,%v) are not valid x and y coordinates\n", initialX, initialY))
	}

	if !initialOrientation.IsValid() {
		return "", errors.New(fmt.Sprintf("%v is not a valid initial orientation\n", initialOrientation))
	}

	r.currentX = initialX
	r.currentY = initialY
	r.currentOrientation = initialOrientation

	for _, v := range commands {
		switch v {
		case Left:
			r.TurnLeft()
		case Right:
			r.TurnRight()
		case Advance:
			err := r.Advance()
			if err != nil {
				return r.formatOutput(false), nil
			}
		}
	}

	return r.formatOutput(true), nil
}

// TurnRight will change Rover's current orientation to the next CardinalPoint clockwise.
func (r *Rover) TurnRight() {

	switch r.currentOrientation {
	case North:
		r.currentOrientation = East
	case East:
		r.currentOrientation = South
	case South:
		r.currentOrientation = West
	case West:
		r.currentOrientation = North
	}
}

// TurnLeft will change Rover's current orientation to the next CardinalPoint counterclockwise.
func (r *Rover) TurnLeft() {

	switch r.currentOrientation {
	case North:
		r.currentOrientation = West
	case West:
		r.currentOrientation = South
	case South:
		r.currentOrientation = East
	case East:
		r.currentOrientation = North
	}
}

// Advance will move the Rover's position adding or subtracting 1 to the actual coordinates based on the currentOrientation.
func (r *Rover) Advance() error {
	newCoordinateX := r.currentX
	newCoordinateY := r.currentY

	switch r.currentOrientation {
	case North:
		newCoordinateY = r.currentY + 1
	case West:
		newCoordinateX = r.currentX - 1
	case South:
		newCoordinateY = r.currentY - 1
	case East:
		newCoordinateX = r.currentX + 1
	}

	if !r.navigationMap.IsValid(newCoordinateX, newCoordinateY) {
		return errors.New(fmt.Sprintf("can not advance to (%v,%v)\n", newCoordinateX, newCoordinateY))
	}

	r.currentX = newCoordinateX
	r.currentY = newCoordinateY

	return nil
}

// convertStringToCommands will convert a string into a list of valid commands Rover commands.
func convertStringToCommands(listOfCommands string) ([]Command, error) {

	if listOfCommands == "" {
		return nil, errors.New(fmt.Sprintf("list of commands is empty\n"))
	}

	commands := strings.Split(listOfCommands, "")
	validatedCommands := make([]Command, 0)

	for _, v := range commands {
		command := Command(v)
		if !command.IsValid() {
			return nil, errors.New(fmt.Sprintf("%v is not a valid command\n", command))
		}

		validatedCommands = append(validatedCommands, command)
	}

	return validatedCommands, nil
}

// formatOutput will format the return string to the requested specification.
//
//	-(True , N, (1,4) when the final destination is within the map's limit.
//	-(False , N, (1,10) when the final destination falls out the map's limit.
func (r *Rover) formatOutput(stillInsideTheMap bool) string {
	return fmt.Sprintf("%v, %v, (%v,%v)", cases.Title(language.English).String(fmt.Sprintf("%v", stillInsideTheMap)), r.currentOrientation, r.currentX, r.currentY)
}
