package main

import (
	"bufio"
	"os"
)

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
	duplicatedItem := getDuplicatedItem(pocket1, pocket2)
	return getItemPriority(duplicatedItem)
}

func getDuplicatedItem(pocket1, pocket2 string) rune {
	items := make(map[rune]int)
	for _, char := range pocket1 {
		items[char] = 1
	}
	for _, char := range pocket2 {
		items[char] |= 2
	}
	for char, value := range items {
		if value > 2 {
			return char
		}
	}
	panic("Well, this is not supposed to happen!")
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
}
