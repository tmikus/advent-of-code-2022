package main

func contains(items *[]int, value int) bool {
	for i := 0; i < len(*items); i++ {
		if (*items)[i] == value {
			return true
		}
	}
	return false
}

func countItems(items *[]int, value int) int {
	count := 0
	for i := 0; i < len(*items); i++ {
		if (*items)[i] == value {
			count++
		}
	}
	return count
}

func findLongestChildFlow(
	valves *[]Valve,
	remainingTime int,
	openValves *[]int,
	visitedValves []int,
	currentValveIndex int,
) int {
	if len(*openValves) == len(*valves) {
		return 0
	}
	if remainingTime < 1 {
		return 0
	}
	currentValve := &(*valves)[currentValveIndex]
	maxChildFlow := 0
	for _, nextValveIndex := range currentValve.childValveIndices {
		childFlow := moveToValve(valves, remainingTime-1, openValves, visitedValves, nextValveIndex)
		if childFlow > maxChildFlow {
			maxChildFlow = childFlow
		}
	}
	return maxChildFlow
}

func moveToValve(
	valves *[]Valve,
	remainingTime int,
	openValves *[]int,
	visitedValves []int,
	currentValveIndex int,
) int {
	if countItems(&visitedValves, currentValveIndex) > 2 {
		return 0
	}
	visitedValves = append(visitedValves, currentValveIndex)
	flowAfterOpening := 0
	if !contains(openValves, currentValveIndex) {
		currentValve := &(*valves)[currentValveIndex]
		if currentValve.flowRate > 0 {
			flowAfterOpening = openValve(valves, remainingTime, openValves, visitedValves, currentValveIndex)
		}
	}
	flowAfterMoving := findLongestChildFlow(valves, remainingTime, openValves, visitedValves, currentValveIndex)
	if flowAfterMoving > flowAfterOpening {
		return flowAfterMoving
	}
	return flowAfterOpening
}

func openValve(
	valves *[]Valve,
	remainingTime int,
	openValves *[]int,
	visitedValves []int,
	currentValveIndex int,
) int {
	if remainingTime < 1 {
		return 0
	}
	currentValve := &(*valves)[currentValveIndex]
	remainingTime -= 1
	producedFlow := currentValve.flowRate * remainingTime
	childOpenValves := *openValves
	childOpenValves = append(childOpenValves, currentValveIndex)
	return producedFlow + findLongestChildFlow(valves, remainingTime, &childOpenValves, visitedValves, currentValveIndex)
}
