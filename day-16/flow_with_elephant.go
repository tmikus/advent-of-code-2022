package main

func findBestValveScoreForElephant(
	valves *[]Valve,
	openValves []int,
	remainingTimeMe int,
	remainingTimeElephant int,
	myValveIndex int,
	elephantValveIndex int,
) int {
	bestScore := 0
	if remainingTimeElephant > 1 {
		for index := 0; index < len(*valves); index++ {
			if contains(&openValves, index) {
				continue
			}
			valve := &(*valves)[index]
			if valve.flowRate == 0 {
				continue
			}
			score := computeValveScore(valve, remainingTimeElephant, elephantValveIndex)
			childOpenValves := openValves
			childOpenValves = append(childOpenValves, index)
			score.score += findBestValveScoreForMe(
				valves,
				childOpenValves,
				remainingTimeMe,
				score.timeAfterTravel,
				myValveIndex,
				index,
			)
			if score.score > bestScore {
				bestScore = score.score
			}
		}
	} else if remainingTimeMe > 1 {
		return findBestValveScoreForMe(
			valves,
			openValves,
			remainingTimeMe,
			remainingTimeElephant,
			myValveIndex,
			elephantValveIndex,
		)
	}
	return bestScore
}

func findBestValveScoreForMe(
	valves *[]Valve,
	openValves []int,
	remainingTimeMe int,
	remainingTimeElephant int,
	myValveIndex int,
	elephantValveIndex int,
) int {
	bestScore := 0
	if remainingTimeMe > 1 {
		for index := 0; index < len(*valves); index++ {
			if contains(&openValves, index) {
				continue
			}
			valve := &(*valves)[index]
			if valve.flowRate == 0 {
				continue
			}
			score := computeValveScore(valve, remainingTimeMe, myValveIndex)
			childOpenValves := openValves
			childOpenValves = append(childOpenValves, index)
			score.score += findBestValveScoreForElephant(
				valves,
				childOpenValves,
				score.timeAfterTravel,
				remainingTimeElephant,
				index,
				elephantValveIndex,
			)
			if score.score > bestScore {
				bestScore = score.score
			}
		}
	} else if remainingTimeElephant > 1 {
		return findBestValveScoreForElephant(
			valves,
			openValves,
			remainingTimeMe,
			remainingTimeElephant,
			myValveIndex,
			elephantValveIndex,
		)
	}
	return bestScore
}

func findBestScoreWithElephant(
	valves *[]Valve,
	remainingTime int,
	startIndex int,
) int {
	return findBestValveScoreForMe(
		valves,
		[]int{},
		remainingTime,
		remainingTime,
		startIndex,
		startIndex,
	)
}
