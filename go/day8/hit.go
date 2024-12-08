package main

type hit struct {
	Origin    *cell
	Dest      *cell
	Direction direction
	Distance  int
}
