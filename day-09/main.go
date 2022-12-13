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

func (vec *Vector2) eq(value Vector2) bool {
	return vec.x == value.x && vec.y == value.y
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

type Node struct {
	position      Vector2
	pastPositions []Vector2
}

func (node *Node) IsTouching(other *Node) bool {
	diff := node.position.subtract(&other.position)
	abs := diff.abs()
	return abs.x <= 1 && abs.y <= 1
}

func newNode() Node {
	return Node{
		position:      newVector2(0, 0),
		pastPositions: []Vector2{newVector2(0, 0)},
	}
}

func (node *Node) SetPosition(value Vector2) {
	node.position = value
	node.pastPositions = append(node.pastPositions, value)
}

type State struct {
	nodes []Node
}

func (state *State) RunInstruction(instruction Instructions) {
	for step := 0; step < instruction.value; step++ {
		head := &state.nodes[0]
		head.position = head.position.add(&instruction.direction)
		previousNode := head
		for nodeIndex := 0; nodeIndex < len(state.nodes); nodeIndex++ {
			node := &state.nodes[nodeIndex]
			state.UpdateNode(node, previousNode)
			previousNode = node
		}
	}
}

func (state *State) UpdateNode(node, prevNode *Node) {
	if node.IsTouching(prevNode) {
		return
	}
	diff := node.position.subtract(&prevNode.position)
	update := newVector2(0, 0)
	if diff.eq(newVector2(2, 0)) {
		update = newVector2(-1, 0)
	} else if diff.eq(newVector2(-2, 0)) {
		update = newVector2(1, 0)
	} else if diff.eq(newVector2(0, 2)) {
		update = newVector2(0, -1)
	} else if diff.eq(newVector2(0, -2)) {
		update = newVector2(0, 1)
	} else {
		x := 0
		y := 0
		if diff.x < 0 {
			x = 1
		} else {
			x = -1
		}
		if diff.y < 0 {
			y = 1
		} else {
			y = -1
		}
		update = newVector2(x, y)
	}
	node.SetPosition(node.position.add(&update))
}

func countUniquePositions(state *State) int {
	positions := make(map[string]bool)
	tail := state.nodes[len(state.nodes)-1]
	for _, position := range tail.pastPositions {
		key := fmt.Sprintf("%v;%v", position.x, position.y)
		positions[key] = true
	}
	return len(positions)
}

func initState(ropeLength int) State {
	state := State{
		nodes: []Node{},
	}
	for index := 0; index < ropeLength; index++ {
		state.nodes = append(state.nodes, newNode())
	}
	return state
}

func runInstructions(instructions []Instructions, ropeLength int) State {
	state := initState(ropeLength)
	for _, instruction := range instructions {
		state.RunInstruction(instruction)
	}
	return state
}

func main() {
	instructions := readInstructions()
	state := runInstructions(instructions, 2)
	println("Part 1 result:", countUniquePositions(&state))

	state2 := runInstructions(instructions, 10)
	println("Part 2 result:", countUniquePositions(&state2))
}
