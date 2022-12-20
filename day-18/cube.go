package main

type Cube struct {
	isFacingOutside bool
	position        Vec3
	back            *Cube
	bottom          *Cube
	front           *Cube
	left            *Cube
	right           *Cube
	top             *Cube
	void            bool
	visited         bool
}

func (c *Cube) CountEmptySides(void bool) int {
	if c.void != void {
		return 0
	}
	result := 0
	if c.top == nil || c.top.void != void {
		result++
	}
	if c.bottom == nil || c.bottom.void != void {
		result++
	}
	if c.back == nil || c.back.void != void {
		result++
	}
	if c.front == nil || c.front.void != void {
		result++
	}
	if c.left == nil || c.left.void != void {
		result++
	}
	if c.right == nil || c.right.void != void {
		result++
	}
	return result
}

func (c *Cube) CountSidesFacingOutside() int {
	if c.void {
		return 0
	}
	result := 0
	if isFacingOutside(c.top) {
		result++
	}
	if isFacingOutside(c.bottom) {
		result++
	}
	if isFacingOutside(c.back) {
		result++
	}
	if isFacingOutside(c.front) {
		result++
	}
	if isFacingOutside(c.left) {
		result++
	}
	if isFacingOutside(c.right) {
		result++
	}
	return result
}

func isFacingOutside(cube *Cube) bool {
	return cube == nil || (cube.void && cube.isFacingOutside)
}

func NewCube(position Vec3) Cube {
	return Cube{
		position: position,
		void:     true,
	}
}
