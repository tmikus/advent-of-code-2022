package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func sum(values []int64, take int) int64 {
	var result int64 = 0
	for index := 0; index < take; index++ {
		result += values[index]
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var currentElfCalories int64 = 0
	elfCalories := make([]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elfCalories = append(elfCalories, currentElfCalories)
			currentElfCalories = 0
			continue
		}
		calories, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal("Expected int32 value, found: ", line)
			return
		}
		currentElfCalories += calories
	}
	if currentElfCalories > 0 {
		elfCalories = append(elfCalories, currentElfCalories)
	}
	sort.Slice(elfCalories, func(i, j int) bool { return elfCalories[i] > elfCalories[j] })
	totalSnacks := sum(elfCalories, 1) // Part 1 answer
	println("Part 1 answer:", totalSnacks)

	totalSnacks = sum(elfCalories, 3) // Part 2 answer
	println("Part 2 answer: ", totalSnacks)
}
