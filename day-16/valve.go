package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Valve struct {
	childValves       []string
	childValveIndices []int
	flowRate          int
	name              string
}

func findValveIndex(valves *[]Valve, name string) int {
	for index := 0; index < len(*valves); index++ {
		if (*valves)[index].name == name {
			return index
		}
	}
	panic("Valve not found")
}

func parseValve(line string) Valve {
	regex, err := regexp.Compile("Valve (.*?) has flow rate=(.*?); tunnels? leads? to valves? (.*)")
	if err != nil {
		panic(fmt.Sprint("Could not compile regex: ", err))
	}
	matches := regex.FindSubmatch([]byte(line))
	name := string(matches[1])
	flowRate := parseInt(string(matches[2]))
	childValves := strings.Split(string(matches[3]), ", ")
	return Valve{
		childValves: childValves,
		flowRate:    flowRate,
		name:        name,
	}
}

func readValves(lines []string) []Valve {
	result := make([]Valve, 0)
	for _, line := range lines {
		result = append(result, parseValve(line))
	}
	updateValveIndices(&result)
	return result
}

func updateValveIndices(valves *[]Valve) {
	for currentValveIndex := 0; currentValveIndex < len(*valves); currentValveIndex++ {
		valve := &(*valves)[currentValveIndex]
		for childValveIndex := 0; childValveIndex < len(valve.childValves); childValveIndex++ {
			valve.childValveIndices = append(
				valve.childValveIndices,
				findValveIndex(valves, valve.childValves[childValveIndex]),
			)
		}
	}
}
