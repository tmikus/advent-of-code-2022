package main

import "math"

func computeShortestDistance(
	valves *[]Valve,
	visitedIndices []int,
	currentIndex int,
	targetIndex int,
) int {
	if currentIndex == targetIndex {
		return 0
	}
	currentValve := &(*valves)[currentIndex]
	visitedIndices = append(visitedIndices, currentIndex)
	minDistance := math.MaxInt
	foundPath := false
	for _, childIndex := range currentValve.childValveIndices {
		if contains(&visitedIndices, childIndex) {
			continue
		}
		distance := computeShortestDistance(valves, visitedIndices, childIndex, targetIndex)
		if distance != -1 && distance < minDistance {
			minDistance = distance
			foundPath = true
		}
	}
	if !foundPath {
		return -1
	}
	return minDistance + 1
}

func findShortestDistance(valves *[]Valve, fromIndex int, toIndex int) int {
	return computeShortestDistance(
		valves,
		[]int{},
		fromIndex,
		toIndex,
	)
}
