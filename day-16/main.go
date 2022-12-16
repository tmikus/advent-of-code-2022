package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findStartIndex(valves *[]Valve) int {
	for index, valve := range *valves {
		if valve.name == "AA" {
			return index
		}
	}
	panic("Start point not found!")
}

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
	startIndex := findStartIndex(&valves)
	println("Part 1 result:", findLongestChildFlow(&valves, 30, startIndex))
	println("Part 2 result:", findBestScoreWithElephant(&valves, 26, startIndex))
}
