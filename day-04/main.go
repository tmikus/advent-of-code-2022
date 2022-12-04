package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	from int
	to   int
}

func areRangesOverlapping(left, right Range) bool {
	return isRangeOverlapping(left, right) || isRangeOverlapping(right, left)
}

func isRangeOverlapping(large, small Range) bool {
	return (large.from <= small.from && large.to >= small.from) || (large.from <= small.to && large.to >= small.to)
}

func areRangesFullyWithinAnother(left, right Range) bool {
	return isRangeFullyWithinAnother(left, right) || isRangeFullyWithinAnother(right, left)
}

func isRangeFullyWithinAnother(large, small Range) bool {
	return large.from <= small.from && large.to >= small.to
}

func rangeFromString(input string) Range {
	values := strings.Split(input, "-")
	from, err := strconv.ParseInt(values[0], 10, 32)
	if err != nil {
		panic("Error")
	}
	to, err := strconv.ParseInt(values[1], 10, 32)
	if err != nil {
		panic("Error")
	}
	return Range{
		from: int(from),
		to:   int(to),
	}
}

func rangeTupleFromLine(input string) (Range, Range) {
	values := strings.Split(input, ",")
	left := rangeFromString(values[0])
	right := rangeFromString(values[1])
	return left, right
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	containedCount := 0
	overlappingCount := 0
	for scanner.Scan() {
		left, right := rangeTupleFromLine(scanner.Text())
		if areRangesFullyWithinAnother(left, right) {
			containedCount += 1
		}
		if areRangesOverlapping(left, right) {
			overlappingCount += 1
		}
	}
	println("Part 1 result:", containedCount)
	println("Part 2 result:", overlappingCount)
}
