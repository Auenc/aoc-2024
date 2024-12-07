package main

type guardMovement struct {
	y    int
	x    int
	char string
}

var (
	up = guardMovement{
		y:    -1,
		x:    0,
		char: "^",
	}
	down = guardMovement{
		y:    1,
		x:    0,
		char: "v",
	}
	left = guardMovement{
		y:    0,
		x:    -1,
		char: "<",
	}
	right = guardMovement{
		y:    0,
		x:    1,
		char: ">",
	}
)

func (g guardMovement) rotate() guardMovement {
	switch g.char {
	case "^":
		return right
	case ">":
		return down
	case "v":
		return left
	case "<":
		return up
	default:
		panic("unknown guard movement")
	}
}
