package main

const X_OFFSET = 2
const Y_OFFSET = 3
const BOARD_WIDTH = 7

type Board struct {
	currentDirectionIndex int
	currentShapeIndex     int
	directions            []MoveDirection
	shapeConstructors     []func(Vec2) Shape
	settledShapes         []Shape
}

func (b *Board) AddShape(shape Shape) {
	b.settledShapes = append(b.settledShapes, shape)
}

func (b *Board) MoveDown(shape *Shape) bool {
	if shape.boundingBox.bottomRight.y == 0 {
		return false
	}
	return b.MoveShape(shape, NewVec2(0, -1))
}

func (b *Board) MoveHorizontally(shape *Shape, direction MoveDirection) bool {
	if direction == Left && shape.boundingBox.topLeft.x == 0 {
		return false
	}
	if direction == Right && shape.boundingBox.bottomRight.x == BOARD_WIDTH-1 {
		return false
	}
	directionVector := getDirectionVector(direction)
	return b.MoveShape(shape, directionVector)
}

func (b *Board) MoveShape(shape *Shape, directionVector Vec2) bool {
	shape.Move(directionVector)
	for shapeIndex := 0; shapeIndex < len(b.settledShapes); shapeIndex++ {
		otherShape := &b.settledShapes[shapeIndex]
		if shape.Intersects(otherShape) { // TODO: This might be slow
			shape.Move(directionVector.Inverse())
			return false
		}
	}
	return true
}

func (b *Board) CreateNewShape() Shape {
	startingPoint := b.GetStartPoint()
	shape := b.shapeConstructors[b.currentShapeIndex](startingPoint)
	b.currentShapeIndex = (b.currentShapeIndex + 1) % len(b.shapeConstructors)
	return shape
}

func (b *Board) GetHighestShape() *Shape {
	var highestShape *Shape
	for index := 0; index < len(b.settledShapes); index++ {
		shape := &b.settledShapes[index]
		if highestShape == nil {
			highestShape = shape
			continue
		}
		if shape.boundingBox.topLeft.y > highestShape.boundingBox.topLeft.y {
			highestShape = shape
			continue
		}
	}
	return highestShape
}

func (b *Board) GetNextDirection() MoveDirection {
	direction := b.directions[b.currentDirectionIndex]
	b.currentDirectionIndex = (b.currentDirectionIndex + 1) % len(b.directions)
	return direction
}

func (b *Board) GetStartPoint() Vec2 {
	highestShape := b.GetHighestShape()
	if highestShape == nil {
		return NewVec2(X_OFFSET, Y_OFFSET)
	}
	return NewVec2(X_OFFSET, highestShape.boundingBox.topLeft.y+Y_OFFSET)
}

func NewBoard(directions []MoveDirection) Board {
	return Board{
		currentDirectionIndex: 0,
		currentShapeIndex:     0,
		directions:            directions,
		shapeConstructors: []func(Vec2) Shape{
			NewMinusShape,
			NewPlusShape,
			NewInvertedLShape,
			NewVerticalLineShape,
			NewBoxShape,
		},
	}
}
