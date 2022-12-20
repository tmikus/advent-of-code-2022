package main

import (
	"bufio"
	"os"
)

func countEmptySides(cubes *[]Cube) int {
	result := 0
	for index := 0; index < len(*cubes); index++ {
		result += (*cubes)[index].CountEmptySides(false)
	}
	return result
}

func readCubePositions() []Vec3 {
	scanner := bufio.NewScanner(os.Stdin)
	cubes := make([]Vec3, 0)
	for scanner.Scan() {
		cubes = append(cubes, parseVec3(scanner.Text()))
	}
	return cubes
}

func main() {
	positions := readCubePositions()
	cubes := createCubeGraph(positions)
	println("Part 1 result:", countEmptySides(&cubes))
}
