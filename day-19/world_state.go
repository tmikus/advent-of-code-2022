package main

import "fmt"

type WorldState struct {
	blueprint          *Blueprint
	clayRobotCount     int
	clayCount          int
	geodeRobotCount    int
	geodeCount         int
	obsidianRobotCount int
	obsidianCount      int
	oreRobotCount      int
	oreCount           int
	remainingTime      int
}

func (s *WorldState) ToString() string {
	result := "World State {\n"
	result += fmt.Sprintf("  clayRobotCount: %v\n", s.clayRobotCount)
	result += fmt.Sprintf("  clayCount: %v\n", s.clayCount)
	result += fmt.Sprintf("  geodeRobotCount: %v\n", s.geodeRobotCount)
	result += fmt.Sprintf("  geodeCount: %v\n", s.geodeCount)
	result += fmt.Sprintf("  obsidianRobotCount: %v\n", s.obsidianRobotCount)
	result += fmt.Sprintf("  obsidianCount: %v\n", s.obsidianCount)
	result += fmt.Sprintf("  oreRobotCount: %v\n", s.oreRobotCount)
	result += fmt.Sprintf("  oreCount: %v\n", s.oreCount)
	result += fmt.Sprintf("  remainingTime: %v\n", s.remainingTime)
	result += "}"
	return result
}

func NewWorldState(blueprint *Blueprint, remainingTime int) WorldState {
	return WorldState{
		blueprint:     blueprint,
		oreRobotCount: 1,
		remainingTime: remainingTime,
	}
}
