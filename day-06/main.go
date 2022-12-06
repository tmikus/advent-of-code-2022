package main

import (
	"bufio"
	"os"
)

const START_OF_PACKET_SIZE = 4
const START_OF_MESSAGE_SIZE = 14

func sliceContainsUniqueValues(slice string) bool {
	seen := make(map[rune]bool)
	for _, char := range slice {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func findPacketStart(line string, sliceSize int) int {
	for index := sliceSize; index < len(line); index++ {
		if !sliceContainsUniqueValues(line[index-sliceSize : index]) {
			continue
		}
		return index
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		result := findPacketStart(line, START_OF_PACKET_SIZE)
		println("Part 1 result:", result)
		result = findPacketStart(line, START_OF_MESSAGE_SIZE)
		println("Part 2 result:", result)
	}
}
