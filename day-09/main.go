package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Instructions struct {
	direction Vector2
	value     int
}

type Vector2 struct {
	x int
	y int
}

func (vec *Vector2) abs() Vector2 {
	return Vector2{
		x: int(math.Abs(float64(vec.x))),
		y: int(math.Abs(float64(vec.y))),
	}
}

func (vec *Vector2) add(value *Vector2) Vector2 {
	return Vector2{
		x: vec.x + value.x,
		y: vec.y + value.y,
	}
}

func (vec *Vector2) subtract(value *Vector2) Vector2 {
	return Vector2{
		x: vec.x - value.x,
		y: vec.y - value.y,
	}
}

func newVector2(x, y int) Vector2 {
	return Vector2{x, y}
}

func parseDirection(value string) Vector2 {
	switch value {
	case "R":
		return newVector2(1, 0)
	case "L":
		return newVector2(-1, 0)
	case "U":
		return newVector2(0, 1)
	case "D":
		return newVector2(0, -1)
	}
	panic(fmt.Sprint("Invalid direction:", value))
}

func parseInstruction(value string) Instructions {
	parts := strings.Split(value, " ")
	return Instructions{
		direction: parseDirection(parts[0]),
		value:     parseInt(parts[1]),
	}
}

func parseInt(value string) int {
	amount, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic(fmt.Sprint("Invalid int:", value, err))
	}
	return int(amount)
}

func readInstructions() []Instructions {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]Instructions, 0)
	for scanner.Scan() {
		instruction := parseInstruction(scanner.Text())
		lines = append(lines, instruction)
	}
	return lines
}

type State struct {
	head              Vector2
	tail              Vector2
	pastTailPositions []Vector2
}

func (state *State) RunInstruction(instruction Instructions) {
	for step := 0; step < instruction.value; step++ {
		state.head = state.head.add(&instruction.direction)
		tailToHeadDistance := state.tail.subtract(&state.head)
		absoluteDistance := tailToHeadDistance.abs()
		if absoluteDistance.x > 1 {
			if absoluteDistance.y > 0 {
				state.SetTail(state.head.subtract(&instruction.direction))
			} else {
				state.SetTail(state.tail.add(&instruction.direction))
			}
		} else if absoluteDistance.y > 1 {
			if absoluteDistance.x > 0 {
				state.SetTail(state.head.subtract(&instruction.direction))
			} else {
				state.SetTail(state.tail.add(&instruction.direction))
			}
		}
	}
}

func (state *State) SetTail(value Vector2) {
	state.tail = value
	state.pastTailPositions = append(state.pastTailPositions, value)
}

func countUniquePositions(state *State) int {
	positions := make(map[string]bool)
	for _, position := range state.pastTailPositions {
		key := fmt.Sprintf("%v;%v", position.x, position.y)
		positions[key] = true
	}
	return len(positions)
}

func initState() State {
	return State{
		head:              newVector2(0, 0),
		tail:              newVector2(0, 0),
		pastTailPositions: []Vector2{newVector2(0, 0)},
	}
}

func runInstructions(instructions []Instructions) State {
	state := initState()
	for _, instruction := range instructions {
		state.RunInstruction(instruction)
	}
	return state
}

func main() {
	instructions := readInstructions()
	state := runInstructions(instructions)
	println("Part 1 result:", countUniquePositions(&state))
}
