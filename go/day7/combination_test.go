package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombination_Calculate(t *testing.T) {
	tests := []struct {
		name     string
		c        Combination
		expected int
	}{
		{
			name:     "add",
			c:        Combination{A: 1, B: 2, Operator: OperatorAdd},
			expected: 3,
		},
		{
			name:     "mult",
			c:        Combination{A: 2, B: 3, Operator: OperatorMult},
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actual := test.c.Calculate()
			assert.Equal(tt, test.expected, actual)
		})
	}
}
