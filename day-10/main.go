package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	cycles int
	noop   bool
	value  int
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
			cycles: 2,
			noop:   false,
			value:  int(value),
		}
	case "noop":
		return Command{
			cycles: 1,
			noop:   true,
			value:  0,
		}
	}
	panic("Unknown command")
}

func printCharacter(cycle, x int) {
	position := (cycle - 1) % 40
	if position >= x-1 && position <= x+1 {
		print("#")
	} else {
		print(" ")
	}
	if (cycle % 40) == 0 {
		print("\n")
	}
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
		for currentCycle := 0; currentCycle < command.cycles; currentCycle++ {
			if contains(EXPECTED_CYCLES, cycle) {
				result += cycle * x
			}
			printCharacter(cycle, x)
			cycle++
		}
		if !command.noop {
			x += command.value
		}
	}
	println("Day 1 result:", result)
}
