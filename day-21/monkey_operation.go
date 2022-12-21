package main

type MonkeyOperation func(left, right int) int

func AddOperation(left, right int) int {
	return left + right
}

func SubtractOperation(left, right int) int {
	return left - right
}

func MultiplyOperation(left, right int) int {
	return left * right
}

func DivideOperation(left, right int) int {
	return left / right
}
