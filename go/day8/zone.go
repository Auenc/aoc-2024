package main

import (
	"fmt"
	"math"
)

type zone struct {
	cells [][]cell
}

func (z zone) GetHits(origin cell) []hit {
	var hits []hit
	hits = append(hits, z.getMatchingAntennaHits(DirectionUp, origin)...)
	hits = append(hits, z.getMatchingAntennaHits(DirectionDown, origin)...)
	hits = append(hits, z.getMatchingAntennaHits(DirectionLeft, origin)...)
	hits = append(hits, z.getMatchingAntennaHits(DirectionRight, origin)...)
	hits = append(hits, z.getMatchingAntennaHits(DirectionDiaUpLeft, origin)...)
	hits = append(hits, z.getMatchingAntennaHits(DirectionDiaUpRight, origin)...)
	hits = append(hits, z.getMatchingAntennaHits(DirectionDiaDownLeft, origin)...)
	hits = append(hits, z.getMatchingAntennaHits(DirectionDiaDownRight, origin)...)
	return hits
}

// This is not actually needed, I misunderstood the problem
func (z zone) getMatchingAntennaHits(dir direction, origin cell) []hit {
	var hits []hit
	for y, checkedY := origin.Y+dir.Y, 0; y >= 0 && y < len(z.cells); y, checkedY = y+dir.Y, checkedY+1 {
		if dir.Y == 0 && checkedY > 0 {
			break
		}
		for x, checkedX := origin.X+dir.X, 0; x >= 0 && x < len(z.cells[y]); x, checkedX = x+dir.X, checkedX+1 {
			if dir.X == 0 && checkedX > 0 {
				break
			}
			c := z.cells[y][x]
			if c.IsAntenna() && c.Frequency == origin.Frequency {
				hits = append(hits, hit{
					Origin:    &origin,
					Dest:      &c,
					Direction: dir,
					Distance:  int(math.Max(math.Abs(float64(origin.X-x)), math.Abs(float64(origin.Y-y)))),
				})
			}
		}
	}
	return hits
}

func (z *zone) findAndMarkAntinode(origin cell) {
	for y := range z.cells {
		for x := range z.cells[y] {
			if x == origin.X && y == origin.Y {
				continue
			}
			if z.cells[y][x].IsAntenna() && z.cells[y][x].Frequency == origin.Frequency {
				diffX := x - origin.X
				diffY := y - origin.Y
				z.markAntiNode(x, y, diffX, diffY, origin)
				z.markAntiNode(x, y, diffX*-1, diffY*-1, origin)
			}
		}
	}
}

func (z zone) markAntiNode(x, y, diffX, diffY int, origin cell) {
	if y+diffY < 0 || y+diffY >= len(z.cells) || x+diffX < 0 || x+diffX >= len(z.cells[y]) {
		return
	}
	z.cells[y+diffY][x+diffX].HasAntinode = true
	z.markAntiNode(x+diffX, y+diffY, diffX, diffY, origin)
}

func (z zone) antinodeFromHit(h hit) *cell {
	antiY := h.Origin.Y + h.Direction.Y*h.Distance
	antiX := h.Origin.X + h.Direction.X*h.Distance
	if antiY < 0 || antiY >= len(z.cells) || antiX < 0 || antiX >= len(z.cells[0]) {
		return nil
	}
	return &z.cells[antiY][antiX]
}

func (z *zone) MarkAntinodes() {
	for y := range z.cells {
		for x := range z.cells[y] {
			if z.cells[y][x].IsAntenna() {
				z.findAndMarkAntinode(z.cells[y][x])
			}
		}
	}
}

func (z zone) CountAntinodes() int {
	count := 0
	for y := range z.cells {
		for x := range z.cells[y] {
			if z.cells[y][x].HasAntinode {
				count++
			}
		}
	}
	return count
}

func (z zone) PrintAntinodes() {
	for y := range z.cells {
		for x := range z.cells[y] {
			if z.cells[y][x].HasAntinode {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (z zone) PrintFrequencies() {
	for y := range z.cells {
		for x := range z.cells[y] {
			fmt.Print(z.cells[y][x].Frequency)
		}
		fmt.Println()
	}
}

func (z zone) PrintAntinodeFreq() {
	for y := range z.cells {
		for x := range z.cells[y] {
			if z.cells[y][x].Frequency != "." {
				fmt.Print(z.cells[y][x].Frequency)
			} else if z.cells[y][x].HasAntinode {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func NewZone(cells [][]cell) zone {
	return zone{cells: cells}
}
