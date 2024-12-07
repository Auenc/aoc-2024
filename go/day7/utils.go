package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	ErrParsingTestValue = errors.New("error parsing test value")
	ErrParsingNumber    = errors.New("error parsing number")
)

func LoadContents(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ParseCalibration(line string) (*Calibration, error) {
	parts := strings.Split(line, ":")
	testValue, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, errors.Join(ErrParsingTestValue, err)
	}

	numbers := strings.Split(parts[1], " ")
	var numbersInt = make([]int, 0)
	for _, n := range numbers {
		if n == "" {
			continue
		}
		parsed, err := strconv.Atoi(n)
		if err != nil {
			return nil, errors.Join(ErrParsingNumber, err)
		}
		numbersInt = append(numbersInt, parsed)
	}
	calibration := NewCalibration(testValue, numbersInt)
	return &calibration, nil
}

func ParseData(data string) ([]Calibration, error) {
	lines := strings.Split(data, "\n")
	var calibrations = make([]Calibration, len(lines))
	for _, line := range lines {
		calibration, err := ParseCalibration(line)
		if err != nil {
			return nil, err
		}
		calibrations = append(calibrations, *calibration)
	}
	return calibrations, nil
}

func Keys(in map[Tuple][]Combination) []Tuple {
	var keys = make([]Tuple, len(in))
	for k := range in {
		keys = append(keys, k)
	}
	return keys
}
