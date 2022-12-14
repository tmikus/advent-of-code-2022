package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	inspectedItems int
	items          []int
	operation      Operation
	test           Test
}

type Operand int

const (
	Add      Operand = 1
	Multiply         = 2
)

type Operation struct {
	operand Operand
	old     bool
	value   int
}

func (op *Operation) apply(old int) int {
	switch op.operand {
	case Add:
		if op.old {
			return old + old
		}
		return old + op.value
	case Multiply:
		if op.old {
			return old * old
		}
		return old * op.value
	default:
		panic("Invalid operand!")
	}
}

type Test struct {
	divisibleBy     int
	targetWhenFalse int
	targetWhenTrue  int
}

func decreaseWorryLevel(worryLevel int) int {
	return int(math.Floor(float64(worryLevel) / 3.0))
}

func getPart1Result(monkeys *[]Monkey) int {
	sortMonkeysDescByInspectedItems(monkeys)
	monkey1 := (*monkeys)[0]
	monkey2 := (*monkeys)[1]
	return monkey1.inspectedItems * monkey2.inspectedItems
}

func parseMonkey(lines []string) Monkey {
	return Monkey{
		inspectedItems: 0,
		items:          parseStartingItems(lines),
		operation:      parseOperation(lines),
		test:           parseTest(lines),
	}
}

func parseMonkeys(linesList [][]string) []Monkey {
	result := make([]Monkey, 0)
	for _, lines := range linesList {
		result = append(result, parseMonkey(lines))
	}
	return result
}

func parseInt(value string) int {
	parsedValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic("Couldn't parse value")
	}
	return int(parsedValue)
}

func parseOperation(lines []string) Operation {
	operationLine := lines[2]
	parts := strings.Split(operationLine, " new = old ")
	parts = strings.Split(parts[1], " ")
	operand := Add
	switch parts[0] {
	case "+":
		operand = Add
	case "*":
		operand = Multiply
	default:
		panic("Invalid operand!")
	}
	old := parts[1] == "old"
	value := 0
	if !old {
		value = parseInt(parts[1])
	}
	return Operation{
		operand: operand,
		old:     old,
		value:   value,
	}
}

func parseStartingItems(lines []string) []int {
	slices := strings.Split(lines[1], ": ")
	numbers := strings.Split(slices[1], ", ")
	result := make([]int, 0)
	for _, str := range numbers {
		result = append(result, parseInt(str))
	}
	return result
}

func parseTest(lines []string) Test {
	divisibleByParts := strings.Split(lines[3], " divisible by ")
	ifTrueParts := strings.Split(lines[4], " throw to monkey ")
	ifFalseParts := strings.Split(lines[5], " throw to monkey ")
	return Test{
		divisibleBy:     parseInt(divisibleByParts[1]),
		targetWhenFalse: parseInt(ifFalseParts[1]),
		targetWhenTrue:  parseInt(ifTrueParts[1]),
	}
}

func readLines() [][]string {
	scanner := bufio.NewScanner(os.Stdin)
	monkeys := make([][]string, 0)
	monkeyLines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			monkeys = append(monkeys, monkeyLines)
			monkeyLines = make([]string, 0)
			continue
		}
		monkeyLines = append(monkeyLines, line)
	}
	return monkeys
}

func runRound(monkeys *[]Monkey) {
	for monkeyIndex := 0; monkeyIndex < len(*monkeys); monkeyIndex++ {
		runRoundForMonkey(&(*monkeys)[monkeyIndex], monkeys)
	}
}

func runRoundForMonkey(currentMonkey *Monkey, monkeys *[]Monkey) {
	for _, worryLevel := range currentMonkey.items {
		worryLevel = currentMonkey.operation.apply(worryLevel)
		worryLevel = decreaseWorryLevel(worryLevel)
		runTest(worryLevel, currentMonkey, monkeys)
	}
	currentMonkey.inspectedItems += len(currentMonkey.items)
	currentMonkey.items = make([]int, 0)
}

func runRounds(monkeys *[]Monkey, rounds int) {
	for round := 0; round < rounds; round++ {
		runRound(monkeys)
	}
}

func runTest(worryLevel int, currentMonkey *Monkey, monkeys *[]Monkey) {
	if worryLevel%currentMonkey.test.divisibleBy == 0 {
		targetMonkey := &(*monkeys)[currentMonkey.test.targetWhenTrue]
		targetMonkey.items = append(targetMonkey.items, worryLevel)
	} else {
		targetMonkey := &(*monkeys)[currentMonkey.test.targetWhenFalse]
		targetMonkey.items = append(targetMonkey.items, worryLevel)
	}
}

func sortMonkeysDescByInspectedItems(monkeys *[]Monkey) {
	sort.SliceStable(*monkeys, func(i, j int) bool {
		return (*monkeys)[i].inspectedItems > (*monkeys)[j].inspectedItems
	})
}

func main() {
	monkeysLines := readLines()
	monkeys := parseMonkeys(monkeysLines)
	for _, monkey := range monkeys {
		fmt.Println(monkey)
	}
	runRounds(&monkeys, 20)
	println("Part 1 result", getPart1Result(&monkeys))
}
