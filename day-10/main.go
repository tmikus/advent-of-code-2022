package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	noop  bool
	value int
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func parseCommand(value string) Command {
	tokens := strings.Split(value, " ")
	switch tokens[0] {
	case "addx":
		value, err := strconv.ParseInt(tokens[1], 10, 32)
		if err != nil {
			panic("Invalid number")
		}
		return Command{
			noop:  false,
			value: int(value),
		}
	case "noop":
		return Command{
			noop:  true,
			value: 0,
		}
	}
	panic("Unknown command")
}

func readCommands() []Command {
	scanner := bufio.NewScanner(os.Stdin)
	commands := make([]Command, 0)
	for scanner.Scan() {
		line := scanner.Text()
		commands = append(commands, parseCommand(line))
	}
	return commands
}

func main() {
	EXPECTED_CYCLES := []int{20, 60, 100, 140, 180, 220}
	commands := readCommands()
	cycle := 1
	x := 1
	result := 0
	for _, command := range commands {
		cycles := 1
		if !command.noop {
			cycles = 2
		}
		for currentCycle := 0; currentCycle < cycles; currentCycle++ {
			if contains(EXPECTED_CYCLES, cycle) {
				result += (cycle * x)
			}
			cycle++
		}
		if !command.noop {
			x += command.value
		}
	}
	println("Day 1 result:", result)
}
