package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadData(t *testing.T) {
	testMap := [][]string{
		{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "#", ".", ".", "^", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
		{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
	}
	area, err := loadArea("../../data/day6/test.txt")
	assert.NoError(t, err)
	assert.Equal(t, testMap, area.cells)
}

func TestMoveUntilObstacle(t *testing.T) {
	area, _ := loadArea("../../data/day6/test.txt")
	finished := area.moveUntilObstacle()
	assert.False(t, finished)
	expected := [][]string{
		{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", "^", ".", ".", ".", ".", "#"},
		{".", ".", ".", ".", "X", ".", ".", ".", ".", "."},
		{".", ".", "#", ".", "X", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", "X", ".", ".", "#", ".", "."},
		{".", ".", ".", ".", "X", ".", ".", ".", ".", "."},
		{".", "#", ".", ".", "X", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
		{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
	}
	assert.Equal(t, expected, area.cells)
}

func TestComplete(t *testing.T) {
	area, _ := loadArea("../../data/day6/test.txt")
	area.complete()
	expected := [][]string{
		{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", "X", "X", "X", "X", "X", "#"},
		{".", ".", ".", ".", "X", ".", ".", ".", "X", "."},
		{".", ".", "#", ".", "X", ".", ".", ".", "X", "."},
		{".", ".", "X", "X", "X", "X", "X", "#", "X", "."},
		{".", ".", "X", ".", "X", ".", "X", ".", "X", "."},
		{".", "#", "X", "X", "X", "X", "X", "X", "X", "."},
		{".", "X", "X", "X", "X", "X", "X", "X", "#", "."},
		{"#", "X", "X", "X", "X", "X", "X", "X", ".", "."},
		{".", ".", ".", ".", ".", ".", "#", "X", ".", "."},
	}
	assert.Equal(t, expected, area.cells)
}

func TestCountWalked(t *testing.T) {
	area, _ := loadArea("../../data/day6/test.txt")
	area.complete()
	assert.Equal(t, 41, area.countWalked())
}

func TestCountPotentialObstacles(t *testing.T) {
	area, _ := loadArea("../../data/day6/test.txt")
	assert.Equal(t, 6, area.countPotentialObstacles())
}
