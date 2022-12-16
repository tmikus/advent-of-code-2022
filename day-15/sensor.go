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
	distance int
	position Vec2
}

func (s *Sensor) IsWithinSensorRange(x, y int) bool {
	distance := s.distance
	distance -= abs(s.position.x - x)
	distance -= abs(s.position.y - y)
	return distance >= 0
}

func (s *Sensor) IsWithinSensorRangeAndUpdateX(x *int, y int) bool {
	distance := s.distance
	distanceWithoutY := distance - abs(s.position.y-y)
	distanceWithoutXAndY := distanceWithoutY - abs(s.position.x-*x)
	if distanceWithoutXAndY < 0 {
		return false
	}
	// If the point is to the left of position
	if *x < s.position.x {
		*x = *x + distanceWithoutXAndY + distanceWithoutY + 1
	} else {
		*x = *x + distanceWithoutXAndY + 1
	}
	return true
}

func (s *Sensor) GetSensorValue(x, y int) SensorValue {
	if !s.IsWithinSensorRange(x, y) {
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

func newSensor(beacon Vec2, position Vec2) Sensor {
	deltaX := abs(position.x - beacon.x)
	deltaY := abs(position.y - beacon.y)
	distance := deltaX + deltaY
	return Sensor{
		beacon:   beacon,
		distance: distance,
		position: position,
	}
}

func ParseSensor(line string) Sensor {
	regex, err := regexp.Compile("[xy]=(-?\\d+)")
	if err != nil {
		panic(fmt.Sprintf("Invalid regex: %v", err))
	}
	matches := regex.FindAllSubmatch([]byte(line), -1)
	return newSensor(
		NewVec2(
			parseInt(string(matches[2][1])),
			parseInt(string(matches[3][1])),
		),
		NewVec2(
			parseInt(string(matches[0][1])),
			parseInt(string(matches[1][1])),
		),
	)
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
