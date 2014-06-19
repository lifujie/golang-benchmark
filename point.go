package main

type Point struct {
	x, y int
}

func (p *Point) Nop() {
}

func (p *Point) NopArgs(x, y, z int) {
}
