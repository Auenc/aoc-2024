package main

type direction struct {
	X, Y int
}

var (
	DirectionUp           = direction{X: 0, Y: -1}
	DirectionDown         = direction{X: 0, Y: 1}
	DirectionLeft         = direction{X: -1, Y: 0}
	DirectionRight        = direction{X: 1, Y: 0}
	DirectionDiaUpLeft    = direction{X: -1, Y: -1}
	DirectionDiaUpRight   = direction{X: 1, Y: -1}
	DirectionDiaDownLeft  = direction{X: -1, Y: 1}
	DirectionDiaDownRight = direction{X: 1, Y: 1}
)
