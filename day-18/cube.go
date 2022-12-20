package main

type Cube struct {
	position Vec3
	back     *Cube
	bottom   *Cube
	front    *Cube
	left     *Cube
	right    *Cube
	top      *Cube
	void     bool
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

func NewCube(position Vec3) Cube {
	return Cube{
		position: position,
		void:     true,
	}
}
