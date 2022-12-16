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
	//printGrid(&grid)
	println("Part 1 result:", countScannedFields(&sensors, &gridDimensions, 2000000))
}
