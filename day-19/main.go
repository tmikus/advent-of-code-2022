package main

import (
	"bufio"
	"fmt"
	"os"
)

func computeBlueprintsScore(blueprints *[]Blueprint, remainingTime int) int {
	totalScore := 0
	for _, blueprint := range *blueprints {
		maxGeodes := computeLargestNumberOfMinedGeodes(&blueprint, remainingTime)
		fmt.Printf("Blueprint %v max geodes %v\n", blueprint.id, maxGeodes)
		totalScore += blueprint.id * maxGeodes
	}
	return totalScore
}

func printBlueprints(blueprints *[]Blueprint) {
	for _, bp := range *blueprints {
		println(bp.ToString())
	}
}

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
	printBlueprints(&blueprints)
	println("Part 1 result:", computeBlueprintsScore(&blueprints, 24))
}
