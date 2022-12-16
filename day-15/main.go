package main

import (
	"bufio"
	"os"
	"strconv"
)

func countScannedFields(sensors *[]Sensor, gridDimensions *GridDimensions, y int) int {
	result := 0
	for x := gridDimensions.min.x; x < gridDimensions.max.x; x++ {
		scanned := false
		containsBeacon := false
		containsSensor := false
		for sensorIndex := 0; sensorIndex < len(*sensors); sensorIndex++ {
			switch (*sensors)[sensorIndex].GetSensorValue(x, y) {
			case BeaconPosition:
				containsBeacon = true
			case Scanned:
				scanned = true
			case SensorPosition:
				containsSensor = true
			}
		}
		if scanned && !containsSensor && !containsBeacon {
			result++
		}
	}
	return result
}

func findTuningFrequency(sensors *[]Sensor, searchDistance int) int {
	y := 0
	x := 0
	for ; y <= searchDistance; y++ {
		for x = 0; x <= searchDistance; {
			if areAllSensorsOutsideRange(sensors, &x, y) {
				return x*4000000 + y
			}
		}
	}
	panic("Not found!")
}

func areAllSensorsOutsideRange(sensors *[]Sensor, x *int, y int) bool {
	for sensorIndex := 0; sensorIndex < len(*sensors); sensorIndex++ {
		if (*sensors)[sensorIndex].IsWithinSensorRangeAndUpdateX(x, y) {
			return false
		}
	}
	return true
}

func parseInt(input string) int {
	value, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic("Invalid number")
	}
	return int(value)
}

func readSensors() []Sensor {
	scanner := bufio.NewScanner(os.Stdin)
	sensors := make([]Sensor, 0)
	for scanner.Scan() {
		sensors = append(sensors, ParseSensor(scanner.Text()))
	}
	return sensors
}

func main() {
	sensors := readSensors()
	gridDimensions := GetGridDimensions(&sensors)
	println("Part 1 result:", countScannedFields(&sensors, &gridDimensions, 2000000))
	println("Part 2 result:", findTuningFrequency(&sensors, 4000000))
}
