package main

type cell struct {
	X, Y        int
	Frequency   string
	HasAntinode bool
}

func (c cell) IsAntenna() bool {
	return c.Frequency != "."
}
