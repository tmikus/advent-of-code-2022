package main

import (
	"bufio"
	"os"
	"strconv"
)

func countVisibleTrees(grid [][]int) int {
	count := 0
	for y := 0; y < len(grid); y++ {
		row := grid[y]
		for x := 0; x < len(row); x++ {
			if isTreeVisible(grid, x, y) {
				count++
			}
		}
	}
	return count
}

func getMostScenicTreeScore(grid [][]int) int {
	mostScenicTreeScore := 0
	for y := 0; y < len(grid); y++ {
		row := grid[y]
		for x := 0; x < len(row); x++ {
			score := getViewRange(grid, x, y)
			if score > mostScenicTreeScore {
				mostScenicTreeScore = score
			}
		}
	}
	return mostScenicTreeScore
}

func getViewRange(grid [][]int, x, y int) int {
	rowCount := len(grid)
	colCount := len(grid[y])
	return getViewRangeX(grid, x, y, 1, colCount) *
		getViewRangeX(grid, x, y, -1, 0) *
		getViewRangeY(grid, x, y, 1, rowCount) *
		getViewRangeY(grid, x, y, -1, 0)
}

func getViewRangeX(
	grid [][]int,
	x int,
	y int,
	deltaX int,
	targetX int,
) int {
	treeHeight := grid[y][x]
	viewRange := 0
	for index := x + deltaX; (deltaX > 0 && index < targetX) || (deltaX < 0 && index >= targetX); index += deltaX {
		viewRange++
		leftTreeHeight := grid[y][index]
		if treeHeight <= leftTreeHeight {
			return viewRange
		}
	}
	return viewRange
}

func getViewRangeY(
	grid [][]int,
	x int,
	y int,
	deltaY int,
	targetY int,
) int {
	treeHeight := grid[y][x]
	viewRange := 0
	for index := y + deltaY; (deltaY > 0 && index < targetY) || (deltaY < 0 && index >= targetY); index += deltaY {
		viewRange++
		leftTreeHeight := grid[index][x]
		if treeHeight <= leftTreeHeight {
			return viewRange
		}
	}
	return viewRange
}

func isTreeVisible(grid [][]int, x, y int) bool {
	rowCount := len(grid)
	colCount := len(grid[y])
	// Edges are always visible
	if x == 0 || y == 0 || x == (colCount-1) || y == (rowCount-1) {
		return true
	}
	return isTreeVisibleX(grid, x, y, 1, colCount) ||
		isTreeVisibleX(grid, x, y, -1, 0) ||
		isTreeVisibleY(grid, x, y, 1, rowCount) ||
		isTreeVisibleY(grid, x, y, -1, 0)
}

func isTreeVisibleX(
	grid [][]int,
	x int,
	y int,
	deltaX int,
	targetX int,
) bool {
	treeHeight := grid[y][x]
	for index := x + deltaX; (deltaX > 0 && index < targetX) || (deltaX < 0 && index >= targetX); index += deltaX {
		leftTreeHeight := grid[y][index]
		if treeHeight <= leftTreeHeight {
			return false
		}
	}
	return true
}

func isTreeVisibleY(
	grid [][]int,
	x int,
	y int,
	deltaY int,
	targetY int,
) bool {
	treeHeight := grid[y][x]
	for index := y + deltaY; (deltaY > 0 && index < targetY) || (deltaY < 0 && index >= targetY); index += deltaY {
		leftTreeHeight := grid[index][x]
		if treeHeight <= leftTreeHeight {
			return false
		}
	}
	return true
}

func parseLine(line string) []int {
	heights := make([]int, 0)
	for index := 0; index < len(line); index++ {
		value, err := strconv.ParseInt(
			string(line[index]),
			10,
			32,
		)
		if err != nil {
			panic("Oh no, parsing failed!")
		}
		heights = append(heights, int(value))
	}
	return heights
}

func readGrid() [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	parsedLines := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := parseLine(line)
		parsedLines = append(parsedLines, parsedLine)
	}
	return parsedLines
}

func main() {
	grid := readGrid()
	println("Part 1 result", countVisibleTrees(grid))
	println("Part 2 result", getMostScenicTreeScore(grid))
}
