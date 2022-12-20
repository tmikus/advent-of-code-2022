package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Blueprint struct {
	id                int
	clayRobotCost     RobotCost
	geodeRobotCost    RobotCost
	obsidianRobotCost RobotCost
	oreRobotCost      RobotCost
}

func (b *Blueprint) ToString() string {
	result := fmt.Sprintf("Blueprint %v:\n", b.id)
	result += fmt.Sprintf("Clay robot: %v\n", b.clayRobotCost.ToString())
	result += fmt.Sprintf("Geode robot: %v\n", b.geodeRobotCost.ToString())
	result += fmt.Sprintf("Obsidian robot: %v\n", b.obsidianRobotCost.ToString())
	result += fmt.Sprintf("Ore robot: %v\n", b.oreRobotCost.ToString())
	return result
}

func parseInt(value string) int {
	amount, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic(fmt.Sprint("Invalid int:", value, err))
	}
	return int(amount)
}

func ParseBlueprint(line string) Blueprint {
	regex, err := regexp.Compile("\\d+")
	if err != nil {
		panic(fmt.Sprintf("Error parsing regex: %v", err))
	}
	matches := regex.FindAllString(line, -1)
	return Blueprint{
		id:                parseInt(matches[0]),
		oreRobotCost:      NewRobotCost(0, 0, parseInt(matches[1])),
		clayRobotCost:     NewRobotCost(0, 0, parseInt(matches[2])),
		obsidianRobotCost: NewRobotCost(parseInt(matches[4]), 0, parseInt(matches[3])),
		geodeRobotCost:    NewRobotCost(0, parseInt(matches[6]), parseInt(matches[5])),
	}
}
