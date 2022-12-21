package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Number struct {
	moved bool
	value int
}

func NewNumber(value int) Number {
	return Number{value: value}
}

func findZeroIndex(numbers *[]Number) int {
	for i := 0; i < len(*numbers); i++ {
		if (*numbers)[i].value == 0 {
			return i
		}
	}
	panic("Number 0 not found!")
}

func findNextNumberIndex(numbers *[]Number) int {
	for i := 0; i < len(*numbers); i++ {
		if !(*numbers)[i].moved {
			return i
		}
	}
	return -1
}

func getGroveNumber(zeroIndex int, numbers *[]Number, offset int) int {
	index := (offset - zeroIndex + 1) % len(*numbers)
	return (*numbers)[index].value
}

func getGroveCoordinates(numbers *[]Number) int {
	zeroIndex := findZeroIndex(numbers)
	return getGroveNumber(zeroIndex, numbers, 1000) + getGroveNumber(zeroIndex, numbers, 2000) + getGroveNumber(zeroIndex, numbers, 3000)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func moveNumbers(numbers *[]Number, fromIndex int, toIndex int, targetIndex int) {
	maxIndex := max(fromIndex, toIndex)
	minIndex := min(fromIndex, toIndex)
	if targetIndex < minIndex {
		for i := 0; i <= maxIndex-minIndex; i++ {
			(*numbers)[targetIndex+i] = (*numbers)[minIndex+i]
		}
	} else {
		for i := maxIndex - minIndex; i >= 0; i-- {
			(*numbers)[targetIndex+i] = (*numbers)[minIndex+i]
		}
	}

}

func moveNextNumber(numbers *[]Number) bool {
	foundIndex := findNextNumberIndex(numbers)
	if foundIndex == -1 {
		return false
	}
	numberToMove := (*numbers)[foundIndex]
	offset := numberToMove.value % len(*numbers)
	if offset < 0 {
		offset -= 1
	}
	targetIndex := (foundIndex + offset) % len(*numbers)
	if targetIndex < 0 {
		targetIndex = len(*numbers) + targetIndex
	}
	if targetIndex < foundIndex {
		targetIndex += 1
	}
	if targetIndex > foundIndex {
		moveNumbers(numbers, foundIndex+1, targetIndex, foundIndex)
	} else if targetIndex < foundIndex {
		moveNumbers(numbers, targetIndex, foundIndex-1, targetIndex+1)
	}
	numberToMove.moved = true
	(*numbers)[targetIndex] = numberToMove
	return true
}

func parseInt(input string) int {
	value, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic("Invalid number")
	}
	return int(value)
}

func printNumbers(numbers *[]Number) {
	numbersStr := ""
	indicatorsStr := ""
	for index := 0; index < len(*numbers); index++ {
		number := &(*numbers)[index]
		numberLength := len(numbersStr)
		numbersStr += fmt.Sprintf("%v", number.value)
		numberLength = len(numbersStr) - numberLength
		numbersStr += ", "
		for i := 0; i < numberLength; i++ {
			if number.moved {
				indicatorsStr += "^"
			} else {
				indicatorsStr += " "
			}
		}
		indicatorsStr += "  "
	}
	println(numbersStr)
	println(indicatorsStr)
}

func readNumbers() []Number {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]Number, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numbers = append(numbers, NewNumber(parseInt(line)))
	}
	return numbers
}

func main() {
	numbers := readNumbers()
	//printNumbers(&numbers)
	for moveNextNumber(&numbers) {
		//printNumbers(&numbers)
		continue
	}
	println("Part 1 result:", getGroveCoordinates(&numbers))
}
