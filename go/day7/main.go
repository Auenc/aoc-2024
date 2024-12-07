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

	var total = 0
	for _, calibration := range calibrations {
		if calibration.Possible() {
			total += calibration.TestValue
		}
	}
	fmt.Println(total)
}
