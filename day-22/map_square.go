package main

type MapSquare struct {
	bottom *MapSquare
	coords Vec2
	isWall bool
	left   *MapSquare
	right  *MapSquare
	top    *MapSquare
}

func NewMapSquare(x, y int, isWall bool) MapSquare {
	return MapSquare{
		coords: NewVec2(x, y),
		isWall: isWall,
	}
}

func ParseMapSquares(fields *[][]MapField) *MapSquare {
	squares := initSquares(fields)
	setValidFieldsOnMap(fields, &squares)
	setSquarePointers(&squares)
	return findFirstSquare(&squares)
}

func findFirstSquare(squares *[][]*MapSquare) *MapSquare {
	for y := 0; y < len(*squares); y++ {
		row := &(*squares)[y]
		for x := 0; x < len(*row); x++ {
			square := (*row)[x]
			if square != nil {
				return square
			}
		}
	}
	panic("Could not find the first square!")
}

func getBottomSquare(squares *[][]*MapSquare, x, y int) *MapSquare {
	for try := 1; try <= len(*squares); try++ {
		rowIndex := (y + try) % len(*squares)
		square := (*squares)[rowIndex][x]
		if square != nil {
			return square
		}
	}
	panic("Could not find any squares to the top!")
}

func getTopSquare(squares *[][]*MapSquare, x, y int) *MapSquare {
	for try := 1; try <= len(*squares); try++ {
		rowIndex := (y - try) % len(*squares)
		if rowIndex < 0 {
			rowIndex = len(*squares) + rowIndex
		}
		square := (*squares)[rowIndex][x]
		if square != nil {
			return square
		}
	}
	panic("Could not find any squares to the top!")
}

func getLeftSquare(squares *[]*MapSquare, x int) *MapSquare {
	for try := 1; try <= len(*squares); try++ {
		nextX := x - try
		square := getSquareByIndex(squares, nextX)
		if square != nil {
			return square
		}
	}
	panic("Could not find any squares to the left!")
}

func getRightSquare(squares *[]*MapSquare, x int) *MapSquare {
	for try := 1; try <= len(*squares); try++ {
		nextX := x + try
		square := getSquareByIndex(squares, nextX)
		if square != nil {
			return square
		}
	}
	panic("Could not find any squares to the right!")
}

func getSquareByIndex(squares *[]*MapSquare, index int) *MapSquare {
	index = index % len(*squares)
	if index < 0 {
		index = len(*squares) + index
	}
	return (*squares)[index]
}

func initSquares(fields *[][]MapField) [][]*MapSquare {
	result := make([][]*MapSquare, 0)
	for y := 0; y < len(*fields); y++ {
		row := &(*fields)[y]
		result = append(result, make([]*MapSquare, len(*row)))
	}
	return result
}

func setSquarePointers(squares *[][]*MapSquare) {
	for y := 0; y < len(*squares); y++ {
		row := &(*squares)[y]
		for x := 0; x < len(*row); x++ {
			square := (*row)[x]
			if square == nil {
				continue
			}
			// Set the pointers
			square.top = getTopSquare(squares, x, y)
			square.bottom = getBottomSquare(squares, x, y)
			square.left = getLeftSquare(row, x)
			square.right = getRightSquare(row, x)
		}
	}
}

func setValidFieldsOnMap(fields *[][]MapField, squares *[][]*MapSquare) {
	for y := 0; y < len(*fields); y++ {
		row := &(*fields)[y]
		for x := 0; x < len(*row); x++ {
			field := (*row)[x]
			if field == None {
				continue
			}
			square := NewMapSquare(x, y, field == Wall)
			(*squares)[y][x] = &square
		}
	}
}
