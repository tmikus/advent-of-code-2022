package main

func printBoard(board *Board) {
	filledSquares := getFilledSquares(board)
	for y := len(filledSquares) - 1; y >= 0; y-- {
		print("|")
		for x := 0; x < BOARD_WIDTH; x++ {
			if filledSquares[y][x] {
				print("#")
			} else {
				print(" ")
			}
		}
		print("|\n")
	}
	println("---------")
}

func getAllPoints(board *Board) []Vec2 {
	points := make([]Vec2, 0)
	for _, shape := range board.settledShapes {
		for _, point := range shape.points {
			points = append(points, point)
		}
	}
	return points
}

func getBoolBoard(height int) [][]bool {
	board := make([][]bool, 0)
	for i := 0; i <= height; i++ {
		board = append(board, make([]bool, BOARD_WIDTH))
	}
	return board
}

func getFilledSquares(board *Board) [][]bool {
	points := getAllPoints(board)
	maxCoords := GetMaxCoords(&points)
	result := getBoolBoard(maxCoords.y)
	for _, point := range points {
		result[point.y][point.x] = true
	}
	return result
}
