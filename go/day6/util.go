package main

import (
	"os"
	"strings"
)

func loadData(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func parseCells(data string) [][]string {
	lines := strings.Split(data, "\n")

	var cells = make([][]string, len(lines))
	for i, line := range lines {
		cells[i] = strings.Split(line, "")
	}
	return cells
}

func loadArea(path string) (area, error) {
	data, err := loadData(path)
	if err != nil {
		return area{}, err
	}
	cells := parseCells(data)
	return area{input: data, cells: cells}, nil
}
