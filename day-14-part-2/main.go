package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Path []Vec2

func parseInt(input string) int {
	value, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic("Invalid number")
	}
	return int(value)
}

func parsePath(input string) Path {
	parts := strings.Split(input, "->")
	path := make([]Vec2, 0)
	for _, part := range parts {
		path = append(path, parseVec2(strings.TrimSpace(part)))
	}
	return path
}

func parseVec2(input string) Vec2 {
	parts := strings.Split(input, ",")
	return NewVec2(parseInt(parts[0]), parseInt(parts[1]))
}

func readPaths() []Path {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]Path, 0)
	for scanner.Scan() {
		lines = append(lines, parsePath(scanner.Text()))
	}
	return lines
}

func main() {
	paths := readPaths()
	grid := NewGrid(paths)
	println("Part 1 result:", AnimateGrainsOfSand(&grid))
}
