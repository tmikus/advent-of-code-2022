package main

import (
	"strconv"
	"strings"
)

type MonkeyDefinition struct {
	leftMonkeyId  string
	id            string
	index         int
	isStatic      bool
	operation     MonkeyOperation
	rightMonkeyId string
	value         int
}

func ParseMonkeyDefinition(line string) MonkeyDefinition {
	parts := strings.Split(line, ": ")
	id := parts[0]
	operationString := parts[1]
	if strings.Contains(operationString, "+") {
		return ParseMonkeyWithOperation(id, operationString, "+", AddOperation)
	}
	if strings.Contains(operationString, "-") {
		return ParseMonkeyWithOperation(id, operationString, "-", SubtractOperation)
	}
	if strings.Contains(operationString, "*") {
		return ParseMonkeyWithOperation(id, operationString, "*", MultiplyOperation)
	}
	if strings.Contains(operationString, "/") {
		return ParseMonkeyWithOperation(id, operationString, "/", DivideOperation)
	}
	return MonkeyDefinition{
		id:       id,
		isStatic: true,
		value:    parseInt(operationString),
	}
}

func ParseMonkeyWithOperation(id string, line string, sign string, operation MonkeyOperation) MonkeyDefinition {
	parts := strings.Split(line, sign)
	return MonkeyDefinition{
		leftMonkeyId:  strings.TrimSpace(parts[0]),
		id:            id,
		isStatic:      false,
		operation:     operation,
		rightMonkeyId: strings.TrimSpace(parts[1]),
	}
}

func parseInt(input string) int {
	value, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic("Invalid number")
	}
	return int(value)
}
