package main

func AnimateGrainsOfSand(grid *Grid) int {
	grainsOfSand := 0
	for {
		hasGrainSettled := animateOneGrainOfSand(grid)
		if !hasGrainSettled {
			return grainsOfSand
		}
		grainsOfSand++
	}
}

func animateOneGrainOfSand(grid *Grid) bool {
	grainPosition := NewVec2(500, 0)
	if grid.IsOccupied(&grainPosition) {
		return false
	}
	for {
		nextPoint := getDownPoint(&grainPosition)
		if grid.IsOutside(&nextPoint) {
			return false
		}
		if !grid.IsOccupied(&nextPoint) {
			grainPosition = nextPoint
			continue
		}
		nextPoint = getLeftPoint(&grainPosition)
		if grid.IsOutside(&nextPoint) {
			return false
		}
		if !grid.IsOccupied(&nextPoint) {
			grainPosition = nextPoint
			continue
		}
		nextPoint = getRightPoint(&grainPosition)
		if grid.IsOutside(&nextPoint) {
			return false
		}
		if !grid.IsOccupied(&nextPoint) {
			grainPosition = nextPoint
			continue
		}
		grid.SetPoint(&grainPosition)
		return true
	}
}

func getDownPoint(current *Vec2) Vec2 {
	return NewVec2(
		current.x,
		current.y+1,
	)
}

func getLeftPoint(current *Vec2) Vec2 {
	return NewVec2(
		current.x-1,
		current.y+1,
	)
}

func getRightPoint(current *Vec2) Vec2 {
	return NewVec2(
		current.x+1,
		current.y+1,
	)
}
