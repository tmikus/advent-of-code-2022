package main

type Vec2 struct {
	x int
	y int
}

func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
}
