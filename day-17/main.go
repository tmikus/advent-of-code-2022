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

func computeBoardHeight(directions []MoveDirection, shapesCount int) int {
	board := NewBoard(directions)
	return simulateShapes(&board, shapesCount)
}

func main() {
	line := readLine()
	directions := parseDirections(line)
	println("Part 1 result: ", computeBoardHeight(directions, 2022))
	println("Part 2 result: ", computeBoardHeight(directions, 1000000000000))
}
