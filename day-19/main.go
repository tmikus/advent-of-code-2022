package main

import (
	"bufio"
	"os"
)

func readBlueprints() []Blueprint {
	scanner := bufio.NewScanner(os.Stdin)
	blueprints := make([]Blueprint, 0)
	for scanner.Scan() {
		line := scanner.Text()
		blueprints = append(blueprints, ParseBlueprint(line))
	}
	return blueprints
}

func main() {
	blueprints := readBlueprints()
	for _, bp := range blueprints {
		println(bp.ToString())
	}
}
