package main

import "fmt"

type RobotCost struct {
	clay     int
	obsidian int
	ore      int
}

func (r *RobotCost) ToString() string {
	result := ""
	if r.clay > 0 {
		result += fmt.Sprintf("%v clay, ", r.clay)
	}
	if r.obsidian > 0 {
		result += fmt.Sprintf("%v obsidian, ", r.obsidian)
	}
	if r.ore > 0 {
		result += fmt.Sprintf("%v ore, ", r.ore)
	}
	return result
}

func NewRobotCost(clay, obsidian, ore int) RobotCost {
	return RobotCost{
		clay:     clay,
		obsidian: obsidian,
		ore:      ore,
	}
}
