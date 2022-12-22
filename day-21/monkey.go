package main

type Monkey interface {
	ComputeUnknownValueOf(id string, otherSide int) int
	DependsOnMonkey(id string) bool
	GetResult() int
}
