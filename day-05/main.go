package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	count int
	from  int
	to    int
}

type Stack struct {
	data []uint8
}

func parseOperation(line string) Operation {
	splitLine := strings.Split(line, " ")
	count, _ := strconv.ParseInt(splitLine[1], 10, 32)
	from, _ := strconv.ParseInt(splitLine[3], 10, 32)
	to, _ := strconv.ParseInt(splitLine[5], 10, 32)
	return Operation{
		count: int(count),
		from:  int(from),
		to:    int(to),
	}
}

func parseStack(lines []string, stackIndex int) Stack {
	data := make([]uint8, 0)
	sliceStartIndex := 4*stackIndex + 1
	// Iterate in reverse to preserve the right order of data on the stack
	for lineIndex := len(lines) - 1; lineIndex >= 0; lineIndex-- {
		line := lines[lineIndex]
		value := line[sliceStartIndex]
		if value == ' ' {
			continue
		}
		data = append(data, value)
	}
	return Stack{data}
}

func parseStacks(lines []string) []Stack {
	firstLine := lines[0]
	stacksCount := len(firstLine) / 4
	stacks := make([]Stack, stacksCount)
	for stackIndex := 0; stackIndex < stacksCount; stackIndex++ {
		stacks = append(stacks, parseStack(lines, stackIndex))
	}
	return stacks
}

func getResult(stacks []Stack) string {
	result := ""
	for _, stack := range stacks {
		result += string(stack.data[len(stack.data)-1])
	}
	return result
}

func readOperations(scanner *bufio.Scanner) []Operation {
	operations := make([]Operation, 0)
	for scanner.Scan() {
		line := scanner.Text()
		operations = append(operations, parseOperation(line))
	}
	return operations
}

func readStacks(scanner *bufio.Scanner) []Stack {
	stackLines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		stackLines = append(stackLines, line)
	}
	stackLinesWithoutLabels := stackLines[:len(stackLines)-1]
	return parseStacks(stackLinesWithoutLabels)
}

func runOperations(stacks []Stack, operations []Operation) []Stack {

	return stacks
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	stacks := readStacks(scanner)
	operations := readOperations(scanner)
	stacks := runOperations(stacks, operations)
	println("Part 1 result:", stacks)
	println("Part 2 result:", operations)
}
