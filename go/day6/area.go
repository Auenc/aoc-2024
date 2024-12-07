package main

type area struct {
	input string
	cells [][]string
}

func (a *area) getGuard() (int, int, guardMovement) {
	for i, row := range a.cells {
		for j, cell := range row {
			if cell == "^" {
				return i, j, up
			}
			if cell == ">" {
				return i, j, right
			}
			if cell == "v" {
				return i, j, down
			}
			if cell == "<" {
				return i, j, left
			}
		}
	}
	return -1, -1, guardMovement{}
}

func (a area) isObstacle(y, x int) bool {
	return a.cells[y][x] == "#" || a.cells[y][x] == "O"
}

func (a *area) reset() {
	a.cells = parseCells(a.input)
}

func (a *area) moveUntilObstacle() bool {
	y, x, direction := a.getGuard()
	if y == -1 || x == -1 {
		panic("guard not found")
	}
	for {
		newY := y + direction.y
		newX := x + direction.x

		if newY < 0 || newY >= len(a.cells) || newX < 0 || newX >= len(a.cells[0]) {
			a.cells[y][x] = "X"
			// fmt.Println("Placing X, leaving board")
			return true
		}

		if a.isObstacle(newY, newX) {
			a.cells[y][x] = direction.char
			// fmt.Println("Placying guard")
			return false
		}
		a.cells[y][x] = "X"
		// fmt.Println("Placing X")
		y = newY
		x = newX
	}
}

func (a *area) rotateGuard() {
	y, x, direction := a.getGuard()
	a.cells[y][x] = direction.rotate().char
}

func (a *area) complete() bool {
	isLoop := false
	maxIterations := 10000
	iterations := 0
	finished := false
	for !finished {
		if iterations > maxIterations {
			isLoop = true
			break
		}
		finished = a.moveUntilObstacle()
		if !finished {
			a.rotateGuard()
		}
		iterations++
	}
	return isLoop
}

func (a area) countWalked() int {
	count := 0
	for _, row := range a.cells {
		for _, cell := range row {
			if cell == "X" {
				count++
			}
		}
	}
	return count
}
