package main

import (
	"bufio"
	"os"
	"strconv"
)

type Number struct {
	after  *Number
	before *Number
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
	println("A:", a)
	println("B:", b)
	println("C:", c)
	return a + b + c
}

func moveNumber(numbers *[]Number, numberToMove *Number) {
	offset := numberToMove.value % (len(*numbers) - 1)
	currentNumber := numberToMove
	if numberToMove.value >= 0 {
		for i := 0; i < offset; i++ {
			currentNumber = currentNumber.after
			// Skip current number
			if currentNumber == numberToMove {
				currentNumber = currentNumber.after
			}
		}
	} else {
		for i := 0; i >= offset; i-- {
			currentNumber = currentNumber.before
			// Skip current number
			if currentNumber == numberToMove {
				currentNumber = currentNumber.before
			}
		}
	}
	if currentNumber == numberToMove {
		return
	}
	// Remove the number from current position
	numberToMove.before.after = numberToMove.after
	numberToMove.after.before = numberToMove.before

	numberToMove.before = currentNumber
	numberToMove.after = currentNumber.after

	currentNumber.after.before = numberToMove
	currentNumber.after = numberToMove
}

func parseInt(input string) int {
	value, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic("Invalid number")
	}
	return int(value)
}

func printPart1Result(numbers []Number) {
	setNumberPointers(&numbers)
	for index := 0; index < len(numbers); index++ {
		number := &numbers[index]
		moveNumber(&numbers, number)
	}
	println("Part 1 result:", getGroveCoordinates(&numbers))
}

func printPart2Result(numbers []Number) {
	setNumberPointers(&numbers)
	// Set decryption key
	for i := 0; i < len(numbers); i++ {
		numbers[i].value *= 811589153
	}
	// Iterate 10 times
	for iteration := 0; iteration < 10; iteration++ {
		for index := 0; index < len(numbers); index++ {
			number := &numbers[index]
			moveNumber(&numbers, number)
		}
	}
	println("Part 2 result:", getGroveCoordinates(&numbers))
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

func main() {
	numbers := readNumbers()
	printPart1Result(numbers)
	printPart2Result(numbers)
}
