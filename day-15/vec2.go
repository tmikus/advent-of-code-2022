package main

type Vec2 struct {
	x int
	y int
}

func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
}

func MaxVec(a *Vec2, b *Vec2) *Vec2 {
	if a.x < b.x {
		return b
	}
	if a.x > b.x {
		return a
	}
	if a.y < b.y {
		return b
	}
	return a
}

func MinVec(a *Vec2, b *Vec2) *Vec2 {
	if a.x < b.x {
		return a
	}
	if a.x > b.x {
		return b
	}
	if a.y < b.y {
		return a
	}
	return b
}
