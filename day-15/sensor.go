package main

import (
	"fmt"
	"regexp"
)

type SensorValue int

const (
	Empty          SensorValue = 0
	BeaconPosition             = 1
	SensorPosition             = 2
	Scanned                    = 3
)

type Sensor struct {
	beacon   Vec2
	position Vec2
}

func (s *Sensor) GetDistance() int {
	deltaX := abs(s.position.x - s.beacon.x)
	deltaY := abs(s.position.y - s.beacon.y)
	return deltaX + deltaY
}

func (s *Sensor) GetSensorValue(x, y int) SensorValue {
	distance := s.GetDistance()
	distance -= abs(s.position.x - x)
	distance -= abs(s.position.y - y)
	if distance < 0 {
		return Empty
	}
	if x == s.position.x && y == s.position.y {
		return SensorPosition
	}
	if x == s.beacon.x && y == s.beacon.y {
		return BeaconPosition
	}
	return Scanned
}

func ParseSensor(line string) Sensor {
	regex, err := regexp.Compile("[xy]=(-?\\d+)")
	if err != nil {
		panic(fmt.Sprintf("Invalid regex: %v", err))
	}
	matches := regex.FindAllSubmatch([]byte(line), -1)
	return Sensor{
		beacon: NewVec2(
			parseInt(string(matches[2][1])),
			parseInt(string(matches[3][1])),
		),
		position: NewVec2(
			parseInt(string(matches[0][1])),
			parseInt(string(matches[1][1])),
		),
	}
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
