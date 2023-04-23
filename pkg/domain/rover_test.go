package rover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoverConstructor(t *testing.T) {
	//Given
	pm := NewMap(4, 4)

	rover := NewRover(pm)

	//When

	//Then
	assert.NotNil(t, rover)
	assert.NotNil(t, rover.navigationMap)
	assert.Equal(t, 0, rover.currentX)
	assert.Equal(t, 0, rover.currentY)
	assert.Equal(t, CardinalPoint(""), rover.currentOrientation)

	//Given
	rover = NewRover(nil)

	//When

	//Then
	assert.Nil(t, rover)
}

func TestCardinalPoint_IsValid(t *testing.T) {

	testCases := []struct {
		name          string
		cardinalPoint CardinalPoint
		asserts       func(cardinalPoint CardinalPoint)
	}{
		{
			name:          "North",
			cardinalPoint: North,
			asserts: func(cardinalPoint CardinalPoint) {
				assert.True(t, cardinalPoint.IsValid())
			},
		}, {
			name:          "East",
			cardinalPoint: East,
			asserts: func(cardinalPoint CardinalPoint) {
				assert.True(t, cardinalPoint.IsValid())
			},
		}, {
			name:          "South",
			cardinalPoint: South,
			asserts: func(cardinalPoint CardinalPoint) {
				assert.True(t, cardinalPoint.IsValid())
			},
		}, {
			name:          "West",
			cardinalPoint: West,
			asserts: func(cardinalPoint CardinalPoint) {
				assert.True(t, cardinalPoint.IsValid())
			},
		}, {
			name:          "Wrong letter",
			cardinalPoint: "A",
			asserts: func(cardinalPoint CardinalPoint) {
				assert.False(t, cardinalPoint.IsValid())
			},
		},
		{
			name:          "Number",
			cardinalPoint: "1",
			asserts: func(cardinalPoint CardinalPoint) {
				assert.False(t, cardinalPoint.IsValid())
			},
		},
		{
			name:          "Long string",
			cardinalPoint: "asd A23",
			asserts: func(cardinalPoint CardinalPoint) {
				assert.False(t, cardinalPoint.IsValid())
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// given
			cp := tt.cardinalPoint

			// when

			//then
			tt.asserts(cp)
		})
	}
}

func TestCommand_IsValid(t *testing.T) {

	testCases := []struct {
		name    string
		command Command
		asserts func(cmd Command)
	}{
		{
			name:    "Advance",
			command: Advance,
			asserts: func(cmd Command) {
				assert.True(t, cmd.IsValid())
			},
		}, {
			name:    "Left",
			command: Left,
			asserts: func(cmd Command) {
				assert.True(t, cmd.IsValid())
			},
		}, {
			name:    "Right",
			command: Right,
			asserts: func(cmd Command) {
				assert.True(t, cmd.IsValid())
			},
		},
		{
			name:    "Wrong command letter",
			command: "B",
			asserts: func(cmd Command) {
				assert.False(t, cmd.IsValid())
			},
		},
		{
			name:    "Wrong command number",
			command: "1",
			asserts: func(cmd Command) {
				assert.False(t, cmd.IsValid())
			},
		},
		{
			name:    "Long string",
			command: "asd A23",
			asserts: func(cmd Command) {
				assert.False(t, cmd.IsValid())
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// given
			cp := tt.command

			// when

			//then
			tt.asserts(cp)
		})
	}
}

func TestConvertStringToCommands(t *testing.T) {

	testCases := []struct {
		name               string
		unverifiedCommands string
		asserts            func(commands []Command, err error)
	}{
		{
			name:               "Valid commands",
			unverifiedCommands: "ALAARA",
			asserts: func(commands []Command, err error) {
				assert.Len(t, commands, 6)
				assert.Nil(t, err)
				assert.Equal(t, Advance, commands[0])
				assert.Equal(t, Left, commands[1])
				assert.Equal(t, Advance, commands[2])
				assert.Equal(t, Advance, commands[3])
				assert.Equal(t, Right, commands[4])
				assert.Equal(t, Advance, commands[5])

			},
		}, {
			name:               "Empty string",
			unverifiedCommands: "",
			asserts: func(commands []Command, err error) {
				assert.Len(t, commands, 0)
				assert.NotNil(t, err)

			},
		}, {
			name:               "Wrong command letter",
			unverifiedCommands: "ALXARA",
			asserts: func(commands []Command, err error) {
				assert.Len(t, commands, 0)
				assert.NotNil(t, err)

			},
		}, {
			name:               "Wrong command number",
			unverifiedCommands: "ALA1RA",
			asserts: func(commands []Command, err error) {
				assert.Len(t, commands, 0)
				assert.NotNil(t, err)

			},
		}, {
			name:               "Wrong command space",
			unverifiedCommands: "ALA RA",
			asserts: func(commands []Command, err error) {
				assert.Len(t, commands, 0)
				assert.NotNil(t, err)

			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// given

			// when
			commands, err := convertStringToCommands(tt.unverifiedCommands)

			//then
			tt.asserts(commands, err)
		})
	}

}

func TestTurnRight(t *testing.T) {
	//Given
	mp := NewPlanetaryMap(7, 7)
	rover := NewRover(*mp)
	rover.currentOrientation = North

	//When
	rover.TurnRight()

	//Then
	assert.Equal(t, East, rover.currentOrientation)

	//When
	rover.TurnRight()

	//Then
	assert.Equal(t, South, rover.currentOrientation)

	//When
	rover.TurnRight()

	//Then
	assert.Equal(t, West, rover.currentOrientation)

	//When
	rover.TurnRight()

	//Then
	assert.Equal(t, North, rover.currentOrientation)

	//When
	rover.currentOrientation = ""

	//Then
	assert.Equal(t, CardinalPoint(""), rover.currentOrientation)

}

func TestTurnLeft(t *testing.T) {
	//Given
	mp := NewPlanetaryMap(7, 7)
	rover := NewRover(*mp)
	rover.currentOrientation = North

	//When
	rover.TurnLeft()

	//Then
	assert.Equal(t, West, rover.currentOrientation)

	//When
	rover.TurnLeft()

	//Then
	assert.Equal(t, South, rover.currentOrientation)

	//When
	rover.TurnLeft()

	//Then
	assert.Equal(t, East, rover.currentOrientation)

	//When
	rover.TurnLeft()

	//Then
	assert.Equal(t, North, rover.currentOrientation)

	//When
	rover.currentOrientation = ""

	//Then
	assert.Equal(t, CardinalPoint(""), rover.currentOrientation)

}

func TestFormatOutput(t *testing.T) {

	pMap := NewPlanetaryMap(6, 6)

	testCases := []struct {
		name                    string
		planetaryMap            *PlanetaryMap
		roverCurrentX           int
		roverCurrentY           int
		roverCurrentOrientation CardinalPoint

		asserts func(output string, rover *Rover)
	}{
		{
			name:                    "Standard output",
			planetaryMap:            pMap,
			roverCurrentX:           3,
			roverCurrentY:           2,
			roverCurrentOrientation: South,

			asserts: func(output string, rover *Rover) {
				assert.Equal(t, "True, S, (3,2)", rover.formatOutput(true))

			},
		}, {
			name:                    "Standard output 2",
			planetaryMap:            pMap,
			roverCurrentX:           6,
			roverCurrentY:           2,
			roverCurrentOrientation: North,

			asserts: func(output string, rover *Rover) {
				assert.Equal(t, "False, N, (6,2)", rover.formatOutput(false))

			},
		}, {
			name:                    "Standard output 3",
			planetaryMap:            pMap,
			roverCurrentX:           10,
			roverCurrentY:           2,
			roverCurrentOrientation: East,

			asserts: func(output string, rover *Rover) {
				assert.Equal(t, "False, E, (10,2)", rover.formatOutput(false))

			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// given
			rv := NewRover(*tt.planetaryMap)
			rv.currentX = tt.roverCurrentX
			rv.currentY = tt.roverCurrentY
			rv.currentOrientation = tt.roverCurrentOrientation

			// when
			formattedOutput := rv.formatOutput(rv.navigationMap.IsValid(rv.currentX, rv.currentY))

			//then
			tt.asserts(formattedOutput, rv)

		})
	}

}

func TestAdvance(t *testing.T) {

	pMap := NewPlanetaryMap(4, 5)

	testCases := []struct {
		name                    string
		roverCurrentX           int
		roverCurrentY           int
		roverCurrentOrientation CardinalPoint
		asserts                 func(err error)
	}{
		{
			name:                    "From bottom left to in 1",
			roverCurrentX:           0,
			roverCurrentY:           0,
			roverCurrentOrientation: East,
			asserts: func(err error) {
				assert.Nil(t, err)
			},
		}, {
			name:                    "From bottom left to in 2",
			roverCurrentX:           0,
			roverCurrentY:           0,
			roverCurrentOrientation: North,
			asserts: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			name:                    "From bottom left to out 1",
			roverCurrentX:           0,
			roverCurrentY:           0,
			roverCurrentOrientation: West,
			asserts: func(err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:                    "From bottom left to out 2",
			roverCurrentX:           0,
			roverCurrentY:           0,
			roverCurrentOrientation: South,
			asserts: func(err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:                    "From bottom right to in 1",
			roverCurrentX:           3,
			roverCurrentY:           0,
			roverCurrentOrientation: West,
			asserts: func(err error) {
				assert.Nil(t, err)
			},
		}, {
			name:                    "From bottom right to in 2",
			roverCurrentX:           3,
			roverCurrentY:           0,
			roverCurrentOrientation: North,
			asserts: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			name:                    "From bottom right to out 1",
			roverCurrentX:           3,
			roverCurrentY:           0,
			roverCurrentOrientation: East,
			asserts: func(err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:                    "From bottom right to out 2",
			roverCurrentX:           3,
			roverCurrentY:           0,
			roverCurrentOrientation: South,
			asserts: func(err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:                    "From top left to in 1",
			roverCurrentX:           0,
			roverCurrentY:           4,
			roverCurrentOrientation: East,
			asserts: func(err error) {
				assert.Nil(t, err)
			},
		}, {
			name:                    "From top left to in 2",
			roverCurrentX:           0,
			roverCurrentY:           4,
			roverCurrentOrientation: South,
			asserts: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			name:                    "From top left to out 1",
			roverCurrentX:           0,
			roverCurrentY:           4,
			roverCurrentOrientation: West,
			asserts: func(err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:                    "From top left to out 2",
			roverCurrentX:           0,
			roverCurrentY:           4,
			roverCurrentOrientation: North,
			asserts: func(err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:                    "From top right to in 1",
			roverCurrentX:           3,
			roverCurrentY:           4,
			roverCurrentOrientation: West,
			asserts: func(err error) {
				assert.Nil(t, err)
			},
		}, {
			name:                    "From top right to in 2",
			roverCurrentX:           3,
			roverCurrentY:           4,
			roverCurrentOrientation: South,
			asserts: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			name:                    "From top right to out 1",
			roverCurrentX:           3,
			roverCurrentY:           4,
			roverCurrentOrientation: East,
			asserts: func(err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:                    "From top right to out 2",
			roverCurrentX:           3,
			roverCurrentY:           4,
			roverCurrentOrientation: North,
			asserts: func(err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// given
			rv := NewRover(*pMap)
			rv.currentX = tt.roverCurrentX
			rv.currentY = tt.roverCurrentY
			rv.currentOrientation = tt.roverCurrentOrientation

			// when
			err := rv.Advance()

			//then
			tt.asserts(err)

		})
	}
}

func TestNavigate(t *testing.T) {

	pMap := NewPlanetaryMap(4, 5)

	testCases := []struct {
		name               string
		initialX           int
		initialY           int
		initialOrientation CardinalPoint
		listOfCommands     string
		asserts            func(string, error)
		planetaryMap       PlanetaryMap
	}{
		{
			name:               "No planetary map present",
			initialX:           0,
			initialY:           0,
			initialOrientation: East,
			listOfCommands:     "",
			planetaryMap:       nil,
			asserts: func(s string, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:               "Empty list of commands",
			initialX:           0,
			initialY:           0,
			initialOrientation: East,
			listOfCommands:     "",
			planetaryMap:       *pMap,
			asserts: func(s string, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:               "Invalid x coordinate",
			initialX:           4,
			initialY:           0,
			initialOrientation: North,
			listOfCommands:     "ALR",
			planetaryMap:       *pMap,
			asserts: func(s string, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:               "Invalid y coordinate",
			initialX:           2,
			initialY:           5,
			initialOrientation: North,
			listOfCommands:     "ALR",
			planetaryMap:       *pMap,
			asserts: func(s string, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:               "Invalid initial orientations",
			initialX:           0,
			initialY:           0,
			initialOrientation: "Wesr",
			listOfCommands:     "ALR",
			planetaryMap:       *pMap,
			asserts: func(output string, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:               "Short trip 1",
			initialX:           0,
			initialY:           0,
			initialOrientation: East,
			listOfCommands:     "AALARA",
			planetaryMap:       *pMap,
			asserts: func(output string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "True, E, (3,1)", output)
			},
		},
		{
			name:               "Short trip 2",
			initialX:           0,
			initialY:           0,
			initialOrientation: East,
			listOfCommands:     "AALAARALA",
			planetaryMap:       *pMap,
			asserts: func(output string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "True, N, (3,3)", output)
			},
		},
		{
			name:               "Long trip 1 out of the map",
			initialX:           0,
			initialY:           0,
			initialOrientation: East,
			listOfCommands:     "AALAARALAAA",
			planetaryMap:       *pMap,
			asserts: func(output string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "False, N, (3,4)", output)
			},
		},
		{
			name:               "Long trip 1 ",
			initialX:           0,
			initialY:           0,
			initialOrientation: East,
			listOfCommands:     "AALAARALAA",
			planetaryMap:       *pMap,
			asserts: func(output string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "True, N, (3,4)", output)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// given
			rv := NewRover(tt.planetaryMap)

			// when
			output, err := rv.Travel(tt.initialX, tt.initialY, tt.initialOrientation, tt.listOfCommands)

			//then
			tt.asserts(output, err)

		})
	}
}
