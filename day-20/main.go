package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Number struct {
	after  *Number
	before *Number
	moved  bool
	value  int
}

func NewNumber(value int) Number {
	return Number{value: value}
}

func findZeroNumber(numbers *[]Number) *Number {
	for i := 0; i < len(*numbers); i++ {
		number := &(*numbers)[i]
		if number.value == 0 {
			return number
		}
	}
	panic("Number 0 not found!")
}

func findNextNumberToMove(numbers *[]Number) *Number {
	for i := 0; i < len(*numbers); i++ {
		number := &(*numbers)[i]
		if !number.moved {
			return number
		}
	}
	return nil
}

func getGroveNumber(zeroNumber *Number, numbers *[]Number, offset int) int {
	index := (offset + 1) % len(*numbers)
	println("Groove number index:", index)
	number := zeroNumber
	for i := 0; i < index; i++ {
		number = number.after
	}
	fmt.Printf("Grove number %v: %v\n", offset, number.value)
	return number.value
}

func getGroveCoordinates(numbers *[]Number) int {
	a := 0
	b := 0
	c := 0
	zeroNumber := findZeroNumber(numbers)
	number := zeroNumber.after
	for i := 1; i <= 3000; i++ {
		if i == 1000 {
			a = number.value
		} else if i == 2000 {
			b = number.value
		} else if i == 3000 {
			c = number.value
		}
		number = number.after
	}
	return a + b + c
	//return getGroveNumber(zeroNumber, numbers, 1000) + getGroveNumber(zeroNumber, numbers, 2000) + getGroveNumber(zeroNumber, numbers, 3000)
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

//     |
// 0, -2, 2, 3

// 0  (before: 3, after: -2)
// -2 (before: 0, after: 2)
// 2  (before: -2, after: 3)
// 3  (before: 2, after: 0)

func moveNumber(numberToMove *Number, offset int) {
	numberToMove.moved = true
	if offset == 0 {
		return
	}
	currentNumber := numberToMove
	if offset > 0 {
		for i := 0; i < offset; i++ {
			currentNumber = currentNumber.after
		}
	} else {
		for i := 0; i >= offset; i-- {
			currentNumber = currentNumber.before
		}
	}
	// Remove the number from current position
	numberToMove.before.after = numberToMove.after
	numberToMove.after.before = numberToMove.before

	numberToMove.before = currentNumber
	numberToMove.after = currentNumber.after

	currentNumber.after.before = numberToMove
	currentNumber.after = numberToMove
}

func moveNextNumber(numbers *[]Number) bool {
	numberToMove := findNextNumberToMove(numbers)
	if numberToMove == nil {
		return false
	}
	offset := numberToMove.value % len(*numbers)
	moveNumber(numberToMove, offset)
	return true
}

func parseInt(input string) int {
	value, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic("Invalid number")
	}
	return int(value)
}

func printNumbersBackward(numbers *[]Number) {
	number := findZeroNumber(numbers)
	numbersStr := ""
	indicatorsStr := ""
	for index := 0; index < len(*numbers); index++ {
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
		number = number.before
	}
	println(numbersStr)
	println(indicatorsStr)
}

func printNumbersForward(numbers *[]Number) {
	number := findZeroNumber(numbers)
	numbersStr := ""
	indicatorsStr := ""
	for index := 0; index < len(*numbers); index++ {
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
		number = number.after
	}
	println(numbersStr)
	println(indicatorsStr)
}

func setNumberPointers(numbers *[]Number) {
	for numberIndex := 0; numberIndex < len(*numbers); numberIndex++ {
		prevNumberIndex := numberIndex - 1
		nextNumberIndex := numberIndex + 1
		if prevNumberIndex < 0 {
			prevNumberIndex = len(*numbers) - 1
		}
		if nextNumberIndex >= len(*numbers) {
			nextNumberIndex = 0
		}
		nextNumber := &(*numbers)[nextNumberIndex]
		number := &(*numbers)[numberIndex]
		prevNumber := &(*numbers)[prevNumberIndex]
		number.after = nextNumber
		number.before = prevNumber
	}
}

func readNumbers() []Number {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]Number, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numbers = append(numbers, NewNumber(parseInt(line)))
	}
	setNumberPointers(&numbers)
	return numbers
}

func main() {
	numbers := readNumbers()
	for moveNextNumber(&numbers) {
		continue
	}
	println("Part 1 result:", getGroveCoordinates(&numbers))
}
