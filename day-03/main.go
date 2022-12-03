package main

import (
	"bufio"
	"math"
	"os"
)

func getBadgePriorities(rucksacks []string) int {
	groupCount := len(rucksacks) / 3
	result := 0
	for groupIndex := 0; groupIndex < groupCount; groupIndex++ {
		startIndex := groupIndex * 3
		endIndex := (groupIndex + 1) * 3
		rucksacksForGroup := rucksacks[startIndex:endIndex]
		result += getGroupPriority(rucksacksForGroup)
	}
	return result
}

func getGroupPriority(rucksacks []string) int {
	commonItem := getCommonItem(rucksacks)
	return getItemPriority(commonItem)
}

func getCommonItem(rucksacks []string) rune {
	items := make(map[rune]int)
	for memberIndex := 0; memberIndex < len(rucksacks); memberIndex++ {
		bit := int(math.Pow(2, float64(memberIndex)))
		for _, char := range rucksacks[memberIndex] {
			items[char] |= bit
		}
	}
	highestValue := 0
	highestChar := ' '
	for char, value := range items {
		if value > highestValue {
			highestValue = value
			highestChar = char
		}
	}
	if highestChar == ' ' {
		panic("Well, this should not have happened...")
	}
	return highestChar
}

func getPriorities(rucksacks []string) int {
	result := 0
	for _, rucksack := range rucksacks {
		result += getPriority(rucksack)
	}
	return result
}

func getPriority(rucksack string) int {
	midPoint := len(rucksack) / 2
	pocket1 := rucksack[:midPoint]
	pocket2 := rucksack[midPoint:]
	duplicatedItem := getCommonItem([]string{pocket1, pocket2})
	return getItemPriority(duplicatedItem)
}

func getItemPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - 'a' + 1
	}
	return int(item) - 'A' + 27
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rucksacks := make([]string, 0)
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}
	println("Part 1 result: ", getPriorities(rucksacks))
	println("Part 2 result: ", getBadgePriorities(rucksacks))
}
