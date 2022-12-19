package main

type Box struct {
	bottomRight Vec2
	height      int
	topLeft     Vec2
}

func (box *Box) Intersects(other *Box) bool {
	return !(other.topLeft.x > box.bottomRight.x ||
		other.bottomRight.x < box.topLeft.x ||
		other.topLeft.y < box.bottomRight.y ||
		other.bottomRight.y > box.topLeft.y)
}

func NewBox(topLeft, bottomRight Vec2) Box {
	height := abs(topLeft.y - bottomRight.y)
	return Box{
		height:      height,
		topLeft:     topLeft,
		bottomRight: bottomRight,
	}
}

func BoxFromPoints(points *[]Vec2) Box {
	topLeft := GetTopLeftCoords(points)
	bottomRight := GetBottomRightCoords(points)
	return NewBox(topLeft, bottomRight)
}

func GetBottomRightCoords(points *[]Vec2) Vec2 {
	firstPoint := &(*points)[0]
	maxX := firstPoint.x
	minY := firstPoint.y
	for index := 1; index < len(*points); index++ {
		point := &(*points)[index]
		if maxX < point.x {
			maxX = point.x
		}
		if minY > point.y {
			minY = point.y
		}
	}
	return NewVec2(maxX, minY)
}

func GetTopLeftCoords(points *[]Vec2) Vec2 {
	firstPoint := &(*points)[0]
	minX := firstPoint.x
	maxY := firstPoint.y
	for index := 1; index < len(*points); index++ {
		point := &(*points)[index]
		if minX > point.x {
			minX = point.x
		}
		if maxY < point.y {
			maxY = point.y
		}
	}
	return NewVec2(minX, maxY)
}
