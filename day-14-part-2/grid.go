package main

type Grid struct {
	grid   [][]bool
	height int
	width  int
}

func NewGrid(paths []Path) Grid {
	dimensions := getMaxDimensions(&paths)
	grid := initGrid(dimensions)
	drawTerrain(&grid, &dimensions)
	drawPaths(&grid, &paths)
	return Grid{
		grid:   grid,
		height: dimensions.y,
		width:  dimensions.x,
	}
}

func (grid *Grid) IsOccupied(point *Vec2) bool {
	if grid.IsOutside(point) {
		return false
	}
	return grid.grid[point.y][point.x]
}

func (grid *Grid) IsOutside(point *Vec2) bool {
	return point.x >= grid.width || point.y >= grid.height || point.x < 0 || point.y < 0
}

func (grid *Grid) SetPoint(point *Vec2) {
	grid.grid[point.y][point.x] = true
}

func drawLine(grid *[][]bool, from *Vec2, to *Vec2) {
	if from.x != to.x {
		drawHorizontalLine(grid, from, to)
	} else {
		drawVerticalLine(grid, from, to)
	}
}

func drawHorizontalLine(grid *[][]bool, from *Vec2, to *Vec2) {
	start := MinVec(from, to)
	end := MaxVec(from, to)
	for index := start.x; index <= end.x; index++ {
		(*grid)[from.y][index] = true
	}
}

func drawTerrain(grid *[][]bool, dimensions *Vec2) {
	from := NewVec2(0, dimensions.y-1)
	to := NewVec2(dimensions.x-1, dimensions.y-1)
	drawLine(grid, &from, &to)
}

func drawVerticalLine(grid *[][]bool, from *Vec2, to *Vec2) {
	start := MinVec(from, to)
	end := MaxVec(from, to)
	for index := start.y; index <= end.y; index++ {
		(*grid)[index][start.x] = true
	}
}

func drawPath(grid *[][]bool, path *Path) {
	for pointIndex := 1; pointIndex < len(*path); pointIndex++ {
		prevPoint := &(*path)[pointIndex-1]
		currentPoint := &(*path)[pointIndex]
		drawLine(grid, prevPoint, currentPoint)
	}
}

func drawPaths(grid *[][]bool, paths *[]Path) {
	for pathIndex := 0; pathIndex < len(*paths); pathIndex++ {
		drawPath(grid, &(*paths)[pathIndex])
	}
}

func getMaxDimensions(paths *[]Path) Vec2 {
	maxX := 0
	maxY := 0
	for _, path := range *paths {
		for _, point := range path {
			if point.x > maxX {
				maxX = point.x
			}
			if point.y > maxY {
				maxY = point.y
			}
		}
	}
	return NewVec2(maxX+1001, maxY+3)
}

func initGrid(dimensions Vec2) [][]bool {
	result := make([][]bool, 0)
	for rowIndex := 0; rowIndex < dimensions.y; rowIndex++ {
		row := make([]bool, dimensions.x)
		result = append(result, row)
	}
	return result
}
