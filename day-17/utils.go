package main

func abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}

func getDirectionVector(direction MoveDirection) Vec2 {
	switch direction {
	case Left:
		return NewVec2(-1, 0)
	case Right:
		return NewVec2(1, 0)
	}
	panic("Invalid direction")
}
