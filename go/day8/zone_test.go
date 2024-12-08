package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZone_getMatchingAntennaHits(t *testing.T) {
	tests := []struct {
		name     string
		zone     zone
		dir      direction
		origin   cell
		expected []hit
	}{
		{
			name: "up 1",
			zone: loadTestZone("..O..\n..O.."),
			dir:  DirectionUp,
			origin: cell{
				X:         2,
				Y:         1,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 1, Frequency: "O"},
					Dest:      &cell{X: 2, Y: 0, Frequency: "O"},
					Direction: DirectionUp,
					Distance:  1,
				},
			},
		},
		{
			name: "up 2",
			zone: loadTestZone("..O..\n.....\n..O.."),
			dir:  DirectionUp,
			origin: cell{
				X:         2,
				Y:         2,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 2, Frequency: "O"},
					Dest:      &cell{X: 2, Y: 0, Frequency: "O"},
					Direction: DirectionUp,
					Distance:  2,
				},
			},
		},
		{
			name: "down 1",
			zone: loadTestZone("..O..\n..O.."),
			dir:  DirectionDown,
			origin: cell{
				X:         2,
				Y:         0,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 0, Frequency: "O"},
					Dest:      &cell{X: 2, Y: 1, Frequency: "O"},
					Direction: DirectionDown,
					Distance:  1,
				},
			},
		},
		{
			name: "down 2",
			zone: loadTestZone("..O..\n....\n..O.."),
			dir:  DirectionDown,
			origin: cell{
				X:         2,
				Y:         0,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 0, Frequency: "O"},
					Dest:      &cell{X: 2, Y: 2, Frequency: "O"},
					Direction: DirectionDown,
					Distance:  2,
				},
			},
		},
		{
			name: "left 1",
			zone: loadTestZone(".OO.."),
			dir:  DirectionLeft,
			origin: cell{
				X:         2,
				Y:         0,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 0, Frequency: "O"},
					Dest:      &cell{X: 1, Y: 0, Frequency: "O"},
					Direction: DirectionLeft,
					Distance:  1,
				},
			},
		},
		{
			name: "left 2",
			zone: loadTestZone("O.O.."),
			dir:  DirectionLeft,
			origin: cell{
				X:         2,
				Y:         0,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 0, Frequency: "O"},
					Dest:      &cell{X: 0, Y: 0, Frequency: "O"},
					Direction: DirectionLeft,
					Distance:  2,
				},
			},
		},
		{
			name: "right 1",
			zone: loadTestZone("..OO."),
			dir:  DirectionRight,
			origin: cell{
				X:         2,
				Y:         0,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 0, Frequency: "O"},
					Dest:      &cell{X: 3, Y: 0, Frequency: "O"},
					Direction: DirectionRight,
					Distance:  1,
				},
			},
		},
		{
			name: "right 2",
			zone: loadTestZone("..O.O"),
			dir:  DirectionRight,
			origin: cell{
				X:         2,
				Y:         0,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 0, Frequency: "O"},
					Dest:      &cell{X: 4, Y: 0, Frequency: "O"},
					Direction: DirectionRight,
					Distance:  2,
				},
			},
		},
		{
			name: "up-left 1",
			zone: loadTestZone(".O...\n..O.."),
			dir:  DirectionDiaUpLeft,
			origin: cell{
				X:         2,
				Y:         1,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 1, Frequency: "O"},
					Dest:      &cell{X: 1, Y: 0, Frequency: "O"},
					Direction: DirectionDiaUpLeft,
					Distance:  1,
				},
			},
		},
		{
			name: "up-left 2",
			zone: loadTestZone(".O...\n.....\n...O."),
			dir:  DirectionDiaUpLeft,
			origin: cell{
				X:         3,
				Y:         2,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 3, Y: 2, Frequency: "O"},
					Dest:      &cell{X: 1, Y: 0, Frequency: "O"},
					Direction: DirectionDiaUpLeft,
					Distance:  2,
				},
			},
		},
		{
			name: "up-right 1",
			zone: loadTestZone("...O.\n..O.."),
			dir:  DirectionDiaUpRight,
			origin: cell{
				X:         2,
				Y:         1,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 1, Frequency: "O"},
					Dest:      &cell{X: 3, Y: 0, Frequency: "O"},
					Direction: DirectionDiaUpRight,
					Distance:  1,
				},
			},
		},
		{
			name: "up-right 2",
			zone: loadTestZone("...O.\n.....\n.O..."),
			dir:  DirectionDiaUpRight,
			origin: cell{
				X:         1,
				Y:         2,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 1, Y: 2, Frequency: "O"},
					Dest:      &cell{X: 3, Y: 0, Frequency: "O"},
					Direction: DirectionDiaUpRight,
					Distance:  2,
				},
			},
		},
		{
			name: "down-right 1",
			zone: loadTestZone("..O..\n...O."),
			dir:  DirectionDiaDownRight,
			origin: cell{
				X:         2,
				Y:         0,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 0, Frequency: "O"},
					Dest:      &cell{X: 3, Y: 1, Frequency: "O"},
					Direction: DirectionDiaDownRight,
					Distance:  1,
				},
			},
		},
		{
			name: "down-right 2",
			zone: loadTestZone("..O..\n.....\n....O"),
			dir:  DirectionDiaDownRight,
			origin: cell{
				X:         2,
				Y:         0,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 0, Frequency: "O"},
					Dest:      &cell{X: 4, Y: 2, Frequency: "O"},
					Direction: DirectionDiaDownRight,
					Distance:  2,
				},
			},
		},
		{
			name: "down-left 1",
			zone: loadTestZone("..O..\n.O..."),
			dir:  DirectionDiaDownLeft,
			origin: cell{
				X:         2,
				Y:         0,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 0, Frequency: "O"},
					Dest:      &cell{X: 1, Y: 1, Frequency: "O"},
					Direction: DirectionDiaDownLeft,
					Distance:  1,
				},
			},
		},
		{
			name: "down-left 2",
			zone: loadTestZone("..O..\n.....\nO...."),
			dir:  DirectionDiaDownLeft,
			origin: cell{
				X:         2,
				Y:         0,
				Frequency: "O",
			},
			expected: []hit{
				{
					Origin:    &cell{X: 2, Y: 0, Frequency: "O"},
					Dest:      &cell{X: 0, Y: 2, Frequency: "O"},
					Direction: DirectionDiaDownLeft,
					Distance:  2,
				},
			},
		},
		{
			name: "offgrid",
			zone: loadTestZone("..O.."),
			dir:  DirectionUp,
			origin: cell{
				X:         2,
				Y:         0,
				Frequency: "O",
			},
			expected: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result := test.zone.getMatchingAntennaHits(test.dir, test.origin)
			assert.EqualValues(tt, test.expected, result)
		})
	}
}

func TestZoneMarkAntinodes(t *testing.T) {
	testMap, err := LoadContents("../../data/day8/test.txt")
	assert.NoError(t, err)
	z := NewZone(ParseMap(testMap))
	z.MarkAntinodes()
	count := z.CountAntinodes()

	assert.Equal(t, 14, count)
}

func loadTestZone(input string) zone {
	cells := ParseMap(input)
	return NewZone(cells)
}
