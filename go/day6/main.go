package main

import (
	"fmt"
	"sync"
)

func main() {
	area, err := loadArea("../../data/day6/data.txt")
	if err != nil {
		panic(err)
	}
	area.complete()
	walked := area.countWalked()
	fmt.Println("part 1", walked)
	area.reset()
	obsticals := countPotentialObstacles(area.input)
	fmt.Println("part 2", obsticals)
}

func countPotentialObstacles(input string) int {
	count := 0
	initialCells := parseCells(input)
	wg := sync.WaitGroup{}
	for y, row := range initialCells {
		for x := range row {
			wg.Add(1)
			go func() {
				defer wg.Done()
				a := area{input: input, cells: parseCells(input)}
				gY, gX, _ := a.getGuard()
				if gY == y && gX == x {
					return
				}
				if a.cells[y][x] == "#" {
					return
				}
				a.cells[y][x] = "O"
				if a.complete() {
					count++
					fmt.Println("Counted ", count)
				}
			}()
		}
	}
	wg.Wait()
	return count
}
