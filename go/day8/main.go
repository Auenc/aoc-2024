package main

import "fmt"

func main() {
	data, err := LoadContents("../../data/day8/data.txt")
	if err != nil {
		panic(err)
	}
	zone := NewZone(ParseMap(data))
	zone.MarkAntinodes()
	zone.PrintAntinodeFreq()
	fmt.Println(zone.CountAntinodes())
}
