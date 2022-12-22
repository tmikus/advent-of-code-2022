package main

import (
	"regexp"
	"strconv"
)

type Movement interface {
	Run(square *MapSquare, facing Facing) (*MapSquare, Facing)
}

type RotateMovement struct {
	rotateLeft bool
}

func (rm RotateMovement) Run(square *MapSquare, facing Facing) (*MapSquare, Facing) {
	if rm.rotateLeft {
		facing = RotateLeft(facing)
	} else {
		facing = RotateRight(facing)
	}
	return square, facing
}

type MoveMovement struct {
	count int
}

func (m MoveMovement) Run(square *MapSquare, facing Facing) (*MapSquare, Facing) {
	for i := 0; i < m.count; i++ {
		var next *MapSquare
		switch facing {
		case Right:
			next = square.right
		case Down:
			next = square.bottom
		case Left:
			next = square.left
		case Up:
			next = square.top
		}
		if next.isWall {
			break
		}
		square = next
	}
	return square, facing
}

func ParseMovements(input string) []Movement {
	regex, _ := regexp.Compile("(\\d+)|(R|L)")
	tokens := regex.FindAllString(input, -1)
	result := make([]Movement, 0)
	for _, token := range tokens {
		result = append(result, parseMovement(token))
	}
	return result
}

func parseInt(input string) int {
	value, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic("Invalid number")
	}
	return int(value)
}

func parseMovement(token string) Movement {
	if token == "L" || token == "R" {
		return RotateMovement{rotateLeft: token == "L"}
	}
	return MoveMovement{count: parseInt(token)}
}
