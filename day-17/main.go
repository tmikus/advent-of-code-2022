package main

import (
	"bufio"
	"os"
)

type MoveDirection int

const (
	Left  MoveDirection = 1
	Right               = 2
)

func parseDirection(char uint8) MoveDirection {
	switch char {
	case '<':
		return Left
	case '>':
		return Right
	}
	panic("Invalid direction")
}

func parseDirections(line string) []MoveDirection {
	directions := make([]MoveDirection, 0)
	for index := 0; index < len(line); index++ {
		directions = append(directions, parseDirection(line[index]))
	}
	return directions
}

func readLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		panic("No input!")
	}
	return scanner.Text()
}

func main() {
	line := readLine()
	directions := parseDirections(line)
	board := NewBoard(directions)
	simulateShapes(&board, 2022)
	height := board.GetHighestShape().boundingBox.topLeft.y + 1
	println("Part 1 result:", height)
}
