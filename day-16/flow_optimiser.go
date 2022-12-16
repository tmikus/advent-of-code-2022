package main

import "fmt"

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

type BestValve struct {
	index int
	score ValveScore
}

func findNextBestValve(
	valves *[]Valve,
	remainingTime int,
	openValves []int,
	currentValveIndex int,
) BestValve {
	bestValveIndex := -1
	bestValveScore := ValveScore{}
	for index := 0; index < len(*valves); index++ {
		if contains(&openValves, index) {
			continue
		}
		valve := &(*valves)[index]
		score := computeValveScore(valve, remainingTime, currentValveIndex)
		if score.score > bestValveScore.score {
			bestValveScore = score
			bestValveIndex = index
		}
	}
	return BestValve{
		index: bestValveIndex,
		score: bestValveScore,
	}
}

func findLongestChildFlow(
	valves *[]Valve,
	remainingTime int,
	currentValveIndex int,
) int {
	score := 0
	var openValves []int
	for {
		if remainingTime <= 0 {
			break
		}
		bestValve := findNextBestValve(valves, remainingTime, openValves, currentValveIndex)
		fmt.Printf("Opening %d, releasing %d score, remaining time %d\n", bestValve.index, bestValve.score.score, bestValve.score.timeAfterTravel)
		score += bestValve.score.score
		remainingTime = bestValve.score.timeAfterTravel
		openValves = append(openValves, bestValve.index)
		currentValveIndex = bestValve.index
	}
	return score
}
