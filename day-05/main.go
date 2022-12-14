package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	count int
	from  int
	to    int
}

type Stack []string

//type Stack struct {
//	data []string
//}

func moveAll(fromStack Stack, toStack Stack, count int) (Stack, Stack) {
	toStack = append(toStack, fromStack[len(fromStack)-count:]...)
	fromStack = fromStack[:len(fromStack)-count]
	return fromStack, toStack
}

func move(fromStack Stack, toStack Stack, count int) (Stack, Stack) {
	for i := 0; i < count; i++ {
		fromStack, toStack = moveOne(fromStack, toStack)
	}
	return fromStack, toStack
}

func moveOne(fromStack Stack, toStack Stack) (Stack, Stack) {
	toStack = append(toStack, fromStack[len(fromStack)-1])
	fromStack = fromStack[:len(fromStack)-1]
	return fromStack, toStack
}

func parseOperation(line string) Operation {
	splitLine := strings.Split(line, " ")
	count, _ := strconv.ParseInt(splitLine[1], 10, 32)
	from, _ := strconv.ParseInt(splitLine[3], 10, 32)
	to, _ := strconv.ParseInt(splitLine[5], 10, 32)
	return Operation{
		count: int(count),
		from:  int(from - 1), // Convert to zero-based index
		to:    int(to - 1),   // Convert to zero-based index
	}
}

func parseStack(lines []string, stackIndex int) Stack {
	data := make([]string, 0)
	sliceStartIndex := 4*stackIndex + 1
	// Iterate in reverse to preserve the right order of data on the stack
	for lineIndex := len(lines) - 1; lineIndex >= 0; lineIndex-- {
		line := lines[lineIndex]
		value := string(line[sliceStartIndex])
		if value == " " {
			continue
		}
		data = append(data, value)
	}
	return data
}

func parseStacks(lines []string) []Stack {
	firstLine := lines[0]
	stacksCount := int(math.Ceil(float64(len(firstLine)) / 4.0))
	stacks := make([]Stack, 0)
	for stackIndex := 0; stackIndex < stacksCount; stackIndex++ {
		stacks = append(stacks, parseStack(lines, stackIndex))
	}
	return stacks
}

func getResult(stacks []Stack) string {
	result := ""
	for _, stack := range stacks {
		result += stack[len(stack)-1]
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

func runOperation(stacks []Stack, operation Operation) []Stack {
	fromStack := stacks[operation.from]
	toStack := stacks[operation.to]
	fromStack, toStack = move(fromStack, toStack, operation.count)
	stacks[operation.from] = fromStack
	stacks[operation.to] = toStack
	return stacks
}

func runOperations(stacks []Stack, operations []Operation) []Stack {
	for _, operation := range operations {
		stacks = runOperation(stacks, operation)
	}
	return stacks
}

func runOptimisedOperation(stacks []Stack, operation Operation) []Stack {
	fromStack := stacks[operation.from]
	toStack := stacks[operation.to]
	fromStack, toStack = moveAll(fromStack, toStack, operation.count)
	stacks[operation.from] = fromStack
	stacks[operation.to] = toStack
	return stacks
}

func runOptimisedOperations(stacks []Stack, operations []Operation) []Stack {
	for _, operation := range operations {
		stacks = runOptimisedOperation(stacks, operation)
	}
	return stacks
}

func cloneStacks(stacks []Stack) []Stack {
	output := make([]Stack, len(stacks))
	for i, stack := range stacks {
		output[i] = make([]string, len(stack))
		copy(output[i], stack)
	}
	return output
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	stacks := readStacks(scanner)
	operations := readOperations(scanner)
	part1Stacks := runOperations(cloneStacks(stacks), operations)
	output := getResult(part1Stacks)
	println("Part 1 result:", output)
	part2Stacks := runOptimisedOperations(cloneStacks(stacks), operations)
	output = getResult(part2Stacks)
	println("Part 2 result:", output)

}
