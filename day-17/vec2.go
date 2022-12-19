package main

type Vec2 struct {
	x int
	y int
}

func (v *Vec2) Add(value *Vec2) *Vec2 {
	v.x += value.x
	v.y += value.y
	return v
}

func (v *Vec2) Equals(other *Vec2) bool {
	return v.x == other.x && v.y == other.y
}

func (v *Vec2) Inverse() Vec2 {
	return NewVec2(v.x*-1, v.y*-1)
}

func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
}

func AreVecTouching(a *Vec2, b *Vec2, axis Axis) bool {
	switch axis {
	case Horizontal:
		return abs(a.x-b.x) <= 1 && a.y == b.y
	case Vertical:
		return abs(a.y-b.y) <= 1 && a.x == b.x
	}
	panic("Invalid axis")
}

func MaxVec(a *Vec2, b *Vec2) *Vec2 {
	if a.x < b.x {
		return b
	}
	if a.x > b.x {
		return a
	}
	if a.y < b.y {
		return b
	}
	return a
}

func MinVec(a *Vec2, b *Vec2) *Vec2 {
	if a.x < b.x {
		return a
	}
	if a.x > b.x {
		return b
	}
	if a.y < b.y {
		return a
	}
	return b
}

func GetMaxCoords(points *[]Vec2) Vec2 {
	firstPoint := &(*points)[0]
	maxX := firstPoint.x
	maxY := firstPoint.y
	for index := 1; index < len(*points); index++ {
		point := &(*points)[index]
		if maxX < point.x {
			maxX = point.x
		}
		if maxY < point.y {
			maxY = point.y
		}
	}
	return NewVec2(maxX, maxY)
}

func GetMinCoords(points *[]Vec2) Vec2 {
	firstPoint := &(*points)[0]
	minX := firstPoint.x
	minY := firstPoint.y
	for index := 1; index < len(*points); index++ {
		point := &(*points)[index]
		if minX > point.x {
			minX = point.x
		}
		if minY < point.y {
			minY = point.y
		}
	}
	return NewVec2(minX, minY)
}
