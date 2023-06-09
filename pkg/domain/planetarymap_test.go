package rover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapConstructor(t *testing.T) {

	pm := NewMap(4, 3)

	assert.NotNil(t, pm, "The new map method returned nil")
	assert.Equal(t, 4, pm.width, "Expected 4 and got %v", pm.width)
	assert.Equal(t, 3, pm.height, "Expected 3 and got %v", pm.height)
}

func TestPlanetaryMapConstructor(t *testing.T) {

	pm := NewPlanetaryMap(4, 3)

	assert.NotNil(t, pm, "The new planetary map method returned nil")

}

func TestIsValid(t *testing.T) {

	ass := assert.New(t)

	testCases := []struct {
		name               string
		planetaryMapWidth  int
		planetaryMapHeight int
		xCoordinate        int
		yCoordinate        int
		expectedResult     bool
		asserts            func(expected, result bool)
	}{
		{
			name:               "Should be In when bottom left corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        0,
			yCoordinate:        0,
			expectedResult:     true,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
		{
			name:               "Should be Out when left of bottom left corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        -1,
			yCoordinate:        0,
			expectedResult:     false,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
		{
			name:               "Should be Out when under of bottom left corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        0,
			yCoordinate:        -1,
			expectedResult:     false,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
		{
			name:               "Should be In when bottom right corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        3,
			yCoordinate:        0,
			expectedResult:     true,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
		{
			name:               "Should be Out when right of bottom right corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        4,
			yCoordinate:        0,
			expectedResult:     false,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
		{
			name:               "Should be Out when under of bottom right corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        3,
			yCoordinate:        -1,
			expectedResult:     false,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
		{
			name:               "Should be In when top right corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        3,
			yCoordinate:        3,
			expectedResult:     true,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
		{
			name:               "Should be Out when right of top right corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        4,
			yCoordinate:        3,
			expectedResult:     false,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
		{
			name:               "Should be Out when above of top right corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        3,
			yCoordinate:        4,
			expectedResult:     false,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
		{
			name:               "Should be In when top left corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        0,
			yCoordinate:        3,
			expectedResult:     true,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
		{
			name:               "Should be Out when left of top left corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        -1,
			yCoordinate:        3,
			expectedResult:     false,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
		{
			name:               "Should be Out when above of top left corner",
			planetaryMapWidth:  4,
			planetaryMapHeight: 4,
			xCoordinate:        0,
			yCoordinate:        4,
			expectedResult:     false,
			asserts: func(expected, result bool) {

				// result
				ass.Equal(expected, result)

			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			// given
			pmPtr := NewPlanetaryMap(tt.planetaryMapWidth, tt.planetaryMapHeight)
			pm := *pmPtr

			// when
			result := pm.IsValid(tt.xCoordinate, tt.yCoordinate)

			//then
			tt.asserts(tt.expectedResult, result)
		})
	}
}
