package main

type GridDimensions struct {
	max Vec2
	min Vec2
}

func GetGridDimensions(sensors *[]Sensor) GridDimensions {
	influenceBox := GetInfluenceBoxForSensor(&(*sensors)[0])
	dimensions := GridDimensions{
		max: influenceBox.max,
		min: influenceBox.min,
	}
	for index := 1; index < len(*sensors); index++ {
		influenceBox = GetInfluenceBoxForSensor(&(*sensors)[index])
		dimensions.max.x = max(dimensions.max.x, influenceBox.max.x)
		dimensions.max.y = max(dimensions.max.y, influenceBox.max.y)
		dimensions.min.x = min(dimensions.min.x, influenceBox.min.x)
		dimensions.min.y = min(dimensions.min.y, influenceBox.min.y)
	}
	return dimensions
}

func max(left, right int) int {
	if left < right {
		return right
	}
	return left
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}
