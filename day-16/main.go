package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInt(input string) int {
	value, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic("Invalid number")
	}
	return int(value)
}

func readLines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	lines := readLines()
	valves := readValves(lines)
	for _, valve := range valves {
		fmt.Printf("%v\n", valve)
	}
	var openValves []int
	println("Part 1 result:", findLongestChildFlow(&valves, 30, &openValves, []int{}, 0))
}
