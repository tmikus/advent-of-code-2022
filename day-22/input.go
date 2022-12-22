package main

import (
	"bufio"
	"os"
)

type Input struct {
	instructions string // TODO: Come up with a better way of representing this
	rawMap       [][]MapField
}

func ReadInput() Input {
	scanner := bufio.NewScanner(os.Stdin)
	rawMap := readMapFields(scanner)
	instructions := readInstructions(scanner)
	return Input{
		instructions: instructions,
		rawMap:       rawMap,
	}
}

func readInstructions(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		panic("Unexpected end of input!")
	}
	return scanner.Text()
}
