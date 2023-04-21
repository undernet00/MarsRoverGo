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
				assert.Equal(t, "True, N, (3,2)", rover.formatOutput(false))

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
