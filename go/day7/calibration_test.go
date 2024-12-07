package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalibration_possibleCombinations(t *testing.T) {
	tests := []struct {
		name     string
		c        Calibration
		expected [][]Combination
	}{
		{
			name: "two nums",
			c: Calibration{
				TestValue: 0,
				Numbers:   []int{1, 2},
			},
			expected: [][]Combination{
				{
					{A: 1, B: 2, Operator: OperatorAdd},
				},
				{
					{A: 1, B: 2, Operator: OperatorMult},
				},
			},
		},
		{
			name: "three nums",
			c: Calibration{
				TestValue: 0,
				Numbers:   []int{1, 2, 3},
			},
			expected: [][]Combination{
				{
					{A: 1, B: 2, Operator: OperatorAdd},
					{A: 2, B: 3, Operator: OperatorAdd},
				},
				{
					{A: 1, B: 2, Operator: OperatorMult},
					{A: 2, B: 3, Operator: OperatorMult},
				},
				{
					{A: 1, B: 2, Operator: OperatorAdd},
					{A: 2, B: 3, Operator: OperatorMult},
				},
				{
					{A: 1, B: 2, Operator: OperatorMult},
					{A: 2, B: 3, Operator: OperatorAdd},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actual := test.c.possibleCombinations()
			assert.ElementsMatch(t, test.expected, actual)
		})
	}
}

func TestCalibration_Possible(t *testing.T) {
	tests := []struct {
		name     string
		c        Calibration
		expected bool
	}{
		{
			name: "190: 10 19",
			c: Calibration{
				TestValue: 190,
				Numbers:   []int{10, 19},
			},
			expected: true,
		},
		{
			name: "3267: 81 40 27",
			c: Calibration{
				TestValue: 3267,
				Numbers:   []int{81, 40, 27},
			},
			expected: true,
		},
		{
			name: "292: 11 6 16 20",
			c: Calibration{
				TestValue: 292,
				Numbers:   []int{11, 6, 16, 20},
			},
			expected: true,
		},
		{
			name: "83: 17 5",
			c: Calibration{
				TestValue: 83,
				Numbers:   []int{17, 5},
			},
			expected: false,
		},
		{
			name: "156: 15 6",
			c: Calibration{
				TestValue: 156,
				Numbers:   []int{15, 6},
			},
			expected: false,
		},
		{
			name: "7290: 6 8 6 15",
			c: Calibration{
				TestValue: 7290,
				Numbers:   []int{6, 8, 6, 15},
			},
			expected: false,
		},
		{
			name: "161011: 16 10 13",
			c: Calibration{
				TestValue: 161011,
				Numbers:   []int{16, 10, 13},
			},
			expected: false,
		},
		{
			name: "192: 17 8 14",
			c: Calibration{
				TestValue: 192,
				Numbers:   []int{17, 8, 14},
			},
			expected: false,
		},
		{
			name: "21037: 9 7 18 13",
			c: Calibration{
				TestValue: 21037,
				Numbers:   []int{9, 7, 18, 13},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actual := test.c.Possible()
			assert.Equal(tt, test.expected, actual)
		})
	}
}
