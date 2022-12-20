package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func computeBlueprintsScore(blueprints *[]Blueprint, remainingTime int) int {
	totalScore := 0
	for _, blueprint := range *blueprints {
		start := time.Now()
		maxGeodes := computeLargestNumberOfMinedGeodes(&blueprint, remainingTime)
		since := time.Since(start)
		fmt.Printf("Blueprint %v max geodes %v. Took %v\n", blueprint.id, maxGeodes, since)
		totalScore += blueprint.id * maxGeodes
	}
	return totalScore
}

func computePart2Result(blueprints *[]Blueprint, remainingTime int) int {
	totalScore := 1
	for index := 0; index < 3; index++ {
		blueprint := &(*blueprints)[index]
		start := time.Now()
		maxGeodes := computeLargestNumberOfMinedGeodes(blueprint, remainingTime)
		since := time.Since(start)
		fmt.Printf("Blueprint %v max geodes %v. Took %v\n", blueprint.id, maxGeodes, since)
		totalScore *= maxGeodes
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
	//printBlueprints(&blueprints)
	println("Part 1 result:", computeBlueprintsScore(&blueprints, 24))
	println("Part 2 result:", computePart2Result(&blueprints, 27))
}
