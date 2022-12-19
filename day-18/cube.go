package main

type Cube struct {
	position Vec3
	back     *Cube
	bottom   *Cube
	front    *Cube
	left     *Cube
	right    *Cube
	top      *Cube
}

func (c *Cube) CountEmptySides() int {
	result := 0
	if c.top == nil {
		result++
	}
	if c.bottom == nil {
		result++
	}
	if c.back == nil {
		result++
	}
	if c.front == nil {
		result++
	}
	if c.left == nil {
		result++
	}
	if c.right == nil {
		result++
	}
	return result
}

func NewCube(position Vec3) Cube {
	return Cube{
		position: position,
	}
}
