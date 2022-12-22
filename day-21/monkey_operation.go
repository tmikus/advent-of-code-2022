package main

type MonkeyOperation int

const (
	AddOperation      MonkeyOperation = 1
	SubtractOperation                 = 2
	MultiplyOperation                 = 3
	DivideOperation                   = 4
)

func ComputeOperation(operation MonkeyOperation, left, right int) int {
	switch operation {
	case AddOperation:
		return left + right
	case SubtractOperation:
		return left - right
	case MultiplyOperation:
		return left * right
	case DivideOperation:
		return left / right
	}
	panic("Invalid operation")
}

// Example operations:
// 150 = x + 300        | x = -150   x = other - right
//
// 150 = x - 300        | x = 450,   x = other + right
//
// 150 = x * 300        | x = 0.5    x = other / right
//
// 150 = x / 300        | x = 45000  x = right * other

func ComputeInverseOperationWithLeftMissing(operation MonkeyOperation, right, other int) int {
	switch operation {
	case AddOperation:
		return other - right
	case SubtractOperation:
		return other + right
	case MultiplyOperation:
		return other / right
	case DivideOperation:
		return right * other
	}
	panic("Invalid operation")
}

// Example operations:
// 150 = 300 + x        | x = -150   x = other - left
//
// 150 = 300 - x        | x = 150    x = (other - left) * -1
//
// 150 = 300 * x        | x = 0.5    x = other / left
//
// 150 = 300 / x        | x = 2      x = left / other

func ComputeInverseOperationWithRightMissing(operation MonkeyOperation, left, other int) int {
	switch operation {
	case AddOperation:
		return other - left
	case SubtractOperation:
		return (other - left) * -1
	case MultiplyOperation:
		return other / left
	case DivideOperation:
		return left / other
	}
	panic("Invalid operation")
}
