package main

import "math"

func simulateShapes(board *Board, shapesToDrop int) int {
	lastShapeIndex := 0
	startShapeIndex := -1
	startHeight := 0
	endShapeIndex := -1
	endHeight := 0
	for ; lastShapeIndex < shapesToDrop; lastShapeIndex++ {
		shape := board.CreateNewShape()
		hasFinishedMoves := simulateShape(board, shape)
		if hasFinishedMoves {
			if startShapeIndex == -1 {
				startShapeIndex = lastShapeIndex
				startHeight = board.GetHighestShape().boundingBox.topLeft.y
			} else {
				endShapeIndex = lastShapeIndex
				endHeight = board.GetHighestShape().boundingBox.topLeft.y
				lastShapeIndex++
				break
			}
		}
	}
	// If we reached the end of the list of shapes before doing pattern-based calculation
	if lastShapeIndex >= shapesToDrop {
		return board.GetHighestShape().boundingBox.topLeft.y + 1
	}
	patternLength := endShapeIndex - startShapeIndex
	countFromStart := shapesToDrop - startShapeIndex
	repetitions := int(math.Floor(float64(countFromStart) / float64(patternLength)))
	repeatedCount := repetitions * patternLength
	remainder := countFromStart - repeatedCount
	currentHeight := startHeight + (repetitions * (endHeight - startHeight))
	for index := 0; index < remainder; index++ {
		shape := board.CreateNewShape()
		hasFinishedMoves := simulateShape(board, shape)
		if hasFinishedMoves {
			if startShapeIndex == -1 {
				startShapeIndex = lastShapeIndex
				startHeight = board.GetHighestShape().boundingBox.topLeft.y
			} else {
				endShapeIndex = lastShapeIndex
				endHeight = board.GetHighestShape().boundingBox.topLeft.y
				lastShapeIndex++
				break
			}
		}
	}
	finalHeight := board.GetHighestShape().boundingBox.topLeft.y - endHeight
	return currentHeight + finalHeight
}

func simulateShape(
	board *Board,
	shape Shape,
) bool {
	hasFinishedMoves := false
	for {
		direction := board.GetNextDirection()
		if board.currentDirectionIndex == 0 {
			hasFinishedMoves = true
		}
		board.MoveHorizontally(&shape, direction)
		if !board.MoveDown(&shape) {
			board.AddShape(shape)
			break
		}
	}
	return hasFinishedMoves
}
