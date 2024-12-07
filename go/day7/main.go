package main

import "fmt"

func main() {
	contents, err := LoadContents("../../data/day7/data.txt")
	if err != nil {
		panic(err)
	}

	calibrations, err := ParseData(contents)
	if err != nil {
		panic(err)
	}

	var correctCalibrations = make([]Calibration, 0)
	var total = 0
	for _, calibration := range calibrations {
		if calibration.Possible() {
			correctCalibrations = append(correctCalibrations, calibration)
			total += calibration.TestValue
		}
	}
	fmt.Println(total)
}
