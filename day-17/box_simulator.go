package main

func simulateShapes(board *Board, shapesToDrop int) {
	for shapeIndex := 0; shapeIndex < shapesToDrop; shapeIndex++ {
		shape := board.CreateNewShape()
		simulateShape(board, shape)
	}
}

func simulateShape(
	board *Board,
	shape Shape,
) {
	for {
		direction := board.GetNextDirection()
		board.MoveHorizontally(&shape, direction)
		if !board.MoveDown(&shape) {
			board.AddShape(shape)
			break
		}
	}
}
