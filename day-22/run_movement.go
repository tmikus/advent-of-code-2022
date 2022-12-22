package main

type Facing int

const (
	Right Facing = 0
	Down         = 1
	Left         = 2
	Up           = 3
)

func GetVec2ForFacing(facing Facing) Vec2 {
	switch facing {
	case Right:
		return NewVec2(1, 0)
	case Down:
		return NewVec2(0, 1)
	case Left:
		return NewVec2(-1, 0)
	case Up:
		return NewVec2(0, -1)
	}
	panic("Invalid facing!")
}

func RotateLeft(facing Facing) Facing {
	switch facing {
	case Right:
		return Up
	case Down:
		return Right
	case Left:
		return Down
	case Up:
		return Left
	}
	panic("Invalid facing!")
}

func RotateRight(facing Facing) Facing {
	switch facing {
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	case Up:
		return Right
	}
	panic("Invalid facing!")
}

func runMovement(startSquare *MapSquare, movements []Movement) (*MapSquare, Facing) {
	facing := Right
	for _, movement := range movements {
		startSquare, facing = movement.Run(startSquare, facing)
	}
	return startSquare, facing
}
