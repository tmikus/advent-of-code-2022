package main

import (
	"bufio"
	"os"
)

type Input struct {
	movements []Movement
	rawMap    [][]MapField
}

func ReadInput() Input {
	scanner := bufio.NewScanner(os.Stdin)
	rawMap := readMapFields(scanner)
	movements := ParseMovements(readInstructions(scanner))
	return Input{
		movements: movements,
		rawMap:    rawMap,
	}
}

func readInstructions(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		panic("Unexpected end of input!")
	}
	return scanner.Text()
}
