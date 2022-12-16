package main

func contains(items *[]int, value int) bool {
	for i := 0; i < len(*items); i++ {
		if (*items)[i] == value {
			return true
		}
	}
	return false
}

type ValveScore struct {
	timeAfterTravel int
	score           int
}

func computeValveScore(
	valve *Valve,
	remainingTime int,
	fromValveIndex int,
) ValveScore {
	distance := valve.distancesFromIndices[fromValveIndex]
	timeAfterTravel := remainingTime - distance - 1
	return ValveScore{
		score:           timeAfterTravel * valve.flowRate,
		timeAfterTravel: timeAfterTravel,
	}
}

func findBestValveScore(
	valves *[]Valve,
	openValves []int,
	remainingTime int,
	currentValveIndex int,
) int {
	if remainingTime < 2 {
		return 0
	}
	bestScore := 0
	for index := 0; index < len(*valves); index++ {
		if contains(&openValves, index) {
			continue
		}
		valve := &(*valves)[index]
		if valve.flowRate == 0 {
			continue
		}
		score := computeValveScore(valve, remainingTime, currentValveIndex)
		childOpenValves := openValves
		childOpenValves = append(childOpenValves, index)
		score.score += findBestValveScore(valves, childOpenValves, score.timeAfterTravel, index)
		if score.score > bestScore {
			bestScore = score.score
		}
	}
	return bestScore
}

func findLongestChildFlow(
	valves *[]Valve,
	remainingTime int,
	currentValveIndex int,
) int {
	return findBestValveScore(
		valves,
		[]int{},
		remainingTime,
		currentValveIndex,
	)
}
