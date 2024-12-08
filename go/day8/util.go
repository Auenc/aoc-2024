package main

import (
	"os"
	"strings"
)

func LoadContents(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ParseMap(data string) [][]cell {
	var rows [][]cell
	for y, line := range strings.Split(data, "\n") {
		var row []cell
		for x, char := range line {
			row = append(row, cell{Frequency: string(char), X: x, Y: y})
		}
		rows = append(rows, row)
	}
	return rows
}
