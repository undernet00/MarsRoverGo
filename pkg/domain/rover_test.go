package rover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoverConstructor(t *testing.T) {
	//Given
	pm := NewPlanetaryMap(4, 4)
	rover := NewRover(pm)

	//When

	//Then
	assert.NotNil(t, rover)
	assert.NotNil(t, rover.navigationMap)
	assert.Equal(t, 0, rover.currentX)
	assert.Equal(t, 0, rover.currentY)
	assert.Equal(t, CardinalPoint(""), rover.currentOrientation)

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
