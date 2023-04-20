package rover

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

type CardinalPoint string

const (
	North CardinalPoint = "N"
	East  CardinalPoint = "E"
	South CardinalPoint = "S"
	West  CardinalPoint = "W"
)

func (cp CardinalPoint) IsValid() bool {
	return cp == North || cp == East || cp == South || cp == West
}

type Command string

const (
	Advance Command = "A"
	Left    Command = "L"
	Right   Command = "R"
)

func (c Command) IsValid(command string) bool {
	anotherCommand := Command(command)
	return anotherCommand == Advance || anotherCommand == Left || anotherCommand == Right
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

func convertStringToCommands(listOfCommands string) ([]Command, error) {

	if listOfCommands == "" {
		return nil, errors.New(fmt.Sprintf("list of commands is empty\n"))
	}

	commands := strings.Split(listOfCommands, "")
	validatedCommands := make([]Command, 0)

	for _, v := range commands {
		command := Command(v)
		if !command.IsValid(v) {
			return nil, errors.New(fmt.Sprintf("%v is not a valid command\n", command))
		}

		validatedCommands = append(validatedCommands, command)
	}

	return validatedCommands, nil
}

func (r *Rover) Traverse(initialX int, initialY int, initialOrientation CardinalPoint, listOfCommands string) (string, error) {

	if r.navigationMap == nil {
		return "", errors.New("not a valid map\n")
	}

	commands, err := convertStringToCommands(listOfCommands)
	if err == nil {
		return "", err
	}

	if !r.navigationMap.IsValid(initialX, initialY) {
		return "", errors.New(spew.Sprintf("(%v,%v) are not valid x and y coordinates\n", initialX, initialY))
	}

	if !initialOrientation.IsValid() {
		return "", errors.New(spew.Sprintf("%v is not a valid initial orientation\n", initialOrientation))
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

func (r *Rover) formatOutput(stillInsideTheMap bool) string {
	return fmt.Sprintf("%v, %v, (%v,%v)", stillInsideTheMap, r.currentOrientation, r.currentX, r.currentY)
}
