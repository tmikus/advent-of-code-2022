package main

import (
	"bufio"
	"os"
)

func readElfPositions() []Vec2 {
	scanner := bufio.NewScanner(os.Stdin)
	positions := make([]Vec2, 0)
	for scanner.Scan() {
		y := len(positions)
		line := scanner.Text()
		for x := 0; x < len(line); x++ {
			if line[x] != '#' {
				continue
			}
			positions = append(positions, NewVec2(x, y))
		}
	}
	return positions
}

func main() {

}
