package main

type SquareType int

const (
	Empty   SquareType = 0
	Falling            = 1
	Settled            = 2
)

func printBoard(board *Board, currentShape *Shape) {
	filledSquares := getFilledSquares(board, currentShape)
	for y := len(filledSquares) - 1; y >= 0; y-- {
		print("|")
		for x := 0; x < BOARD_WIDTH; x++ {
			if filledSquares[y][x] == Settled {
				print("#")
			} else if filledSquares[y][x] == Falling {
				print("@")
			} else {
				print(".")
			}
		}
		print("|\n")
	}
	println("---------")
	print("\n\n")
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

func getEmptyBoardState(height int) [][]SquareType {
	board := make([][]SquareType, 0)
	for i := 0; i <= height; i++ {
		board = append(board, make([]SquareType, BOARD_WIDTH))
	}
	return board
}

func getBoardHeight(points []Vec2, currentShape *Shape) int {
	if currentShape != nil {
		return currentShape.boundingBox.topLeft.y
	}
	maxCoords := GetMaxCoords(&points)
	return maxCoords.y
}

func getFilledSquares(board *Board, currentShape *Shape) [][]SquareType {
	points := getAllPoints(board)
	height := getBoardHeight(points, currentShape)
	result := getEmptyBoardState(height)
	for _, point := range points {
		result[point.y][point.x] = Settled
	}
	if currentShape != nil {
		for _, point := range currentShape.points {
			result[point.y][point.x] = Falling
		}
	}
	return result
}
