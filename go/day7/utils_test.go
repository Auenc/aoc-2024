package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCalibration(t *testing.T) {
	line := "1: 1 2 3"
	calibration, err := ParseCalibration(line)
	assert.NoError(t, err)
	assert.Equal(t, 1, calibration.TestValue)
	assert.Equal(t, []int{1, 2, 3}, calibration.Numbers)
}
