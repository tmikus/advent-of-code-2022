package main

import "math"

func computeLargestNumberOfMinedGeodes(bp *Blueprint, remainingTime int) int {
	state := NewWorldState(bp, remainingTime)
	return constructBestRobot(&state)
}

func constructBestRobot(state *WorldState) int {
	bestScore := 0
	score := constructGeodeRobot(*state)
	if score > bestScore {
		bestScore = score
	}
	score = constructObsidianRobot(*state)
	if score > bestScore {
		bestScore = score
	}
	score = constructClayRobot(*state)
	if score > bestScore {
		bestScore = score
	}
	score = constructOreRobot(*state)
	if score > bestScore {
		bestScore = score
	}
	return bestScore
}

func constructClayRobot(state WorldState) int {
	if !waitUntilCanAfford(&state, &state.blueprint.clayRobotCost) {
		//if state.geodeCount == 11 {
		//	println(state.ToString())
		//}
		return state.geodeCount
	}
	progressTime(&state, 1)
	state.clayRobotCount++
	subtractCost(&state, &state.blueprint.clayRobotCost)
	return constructBestRobot(&state)
}

func constructGeodeRobot(state WorldState) int {
	if !waitUntilCanAfford(&state, &state.blueprint.geodeRobotCost) {
		//if state.geodeCount == 11 {
		//	println(state.ToString())
		//}
		return state.geodeCount
	}
	progressTime(&state, 1)
	state.geodeRobotCount++
	subtractCost(&state, &state.blueprint.geodeRobotCost)
	return constructBestRobot(&state)
}

func constructObsidianRobot(state WorldState) int {
	if !waitUntilCanAfford(&state, &state.blueprint.obsidianRobotCost) {
		//if state.geodeCount == 11 {
		//	println(state.ToString())
		//}
		return state.geodeCount
	}
	progressTime(&state, 1)
	state.obsidianRobotCount++
	subtractCost(&state, &state.blueprint.obsidianRobotCost)
	return constructBestRobot(&state)
}

func constructOreRobot(state WorldState) int {
	if !waitUntilCanAfford(&state, &state.blueprint.oreRobotCost) {
		//if state.geodeCount == 11 {
		//	println(state.ToString())
		//}
		return state.geodeCount
	}
	progressTime(&state, 1)
	state.oreRobotCount++
	subtractCost(&state, &state.blueprint.oreRobotCost)
	return constructBestRobot(&state)
}

func getTimeToAfford(required int, current int, perTurn int) int {
	if required == 0 || required <= current {
		return 0
	}
	if perTurn == 0 {
		panic("It will never finish!")
	}
	return int(math.Ceil(float64(required-current) / float64(perTurn)))
}

func subtractCost(state *WorldState, cost *RobotCost) {
	state.clayCount -= cost.clay
	state.obsidianCount -= cost.obsidian
	state.oreCount -= cost.ore
}

func progressTime(state *WorldState, turns int) {
	state.oreCount += state.oreRobotCount * turns
	state.clayCount += state.clayRobotCount * turns
	state.obsidianCount += state.obsidianRobotCount * turns
	state.geodeCount += state.geodeRobotCount * turns
	state.remainingTime -= turns
}

func waitUntilCanAfford(state *WorldState, cost *RobotCost) bool {
	if cost.clay > 0 && state.clayRobotCount == 0 {
		return false
	}
	if cost.obsidian > 0 && state.obsidianRobotCount == 0 {
		return false
	}
	// Ore always is positive so no problems here
	oreTurns := getTimeToAfford(
		cost.ore,
		state.oreCount,
		state.oreRobotCount,
	)
	clayTurns := getTimeToAfford(
		cost.clay,
		state.clayCount,
		state.clayRobotCount,
	)
	obsidianTurns := getTimeToAfford(
		cost.obsidian,
		state.obsidianCount,
		state.obsidianRobotCount,
	)
	maxTurns := max(oreTurns, clayTurns, obsidianTurns)
	progressTime(state, min(maxTurns, state.remainingTime))
	if maxTurns >= state.remainingTime-1 { // There's no point to build a robot that doesn't improve the result
		return false
	}
	return true
}

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
