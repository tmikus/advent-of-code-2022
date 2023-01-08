package main

type Shape struct {
	boundingBox Box
	points      []Vec2
}

func (s *Shape) Intersects(other *Shape) bool {
	if !s.IsNearTo(other) {
		return false
	}
	if !s.boundingBox.Intersects(&other.boundingBox) {
		return false
	}
	return s.PointsIntersect(other)
}

func (s *Shape) IsNearTo(other *Shape) bool {
	distance := s.boundingBox.topLeft.y - other.boundingBox.topLeft.y
	return distance > -5 || distance < 5
}

func (s *Shape) Move(offset Vec2) {
	s.boundingBox.topLeft.Add(&offset)
	s.boundingBox.bottomRight.Add(&offset)
	for index := 0; index < len(s.points); index++ {
		s.points[index].Add(&offset)
	}
}

func (s *Shape) PointsIntersect(other *Shape) bool {
	for _, point := range s.points {
		for _, otherPoint := range other.points {
			if point.Equals(&otherPoint) {
				return true
			}
		}
	}
	return false
}

func NewShape(points []Vec2) Shape {
	boundingBox := BoxFromPoints(&points)
	return Shape{
		boundingBox: boundingBox,
		points:      points,
	}
}

func NewMinusShape(position Vec2) Shape {
	return NewShape([]Vec2{
		NewVec2(position.x, position.y),   // [#]# # #
		NewVec2(position.x+1, position.y), //  #[#]# #
		NewVec2(position.x+2, position.y), //  # #[#]#
		NewVec2(position.x+3, position.y), //  # # #[#]
	})
}

func NewPlusShape(position Vec2) Shape {
	return NewShape([]Vec2{
		//   [#]
		//  # # #
		//    #
		NewVec2(position.x+1, position.y+2),
		//    #
		// [#]# #
		//    #
		NewVec2(position.x, position.y+1),
		//    #
		//  #[#]#
		//    #
		NewVec2(position.x+1, position.y+1),
		//    #
		//  # #[#]
		//    #
		NewVec2(position.x+2, position.y+1),
		//    #
		//  # # #
		//   [#]
		NewVec2(position.x+1, position.y),
	})
}

func NewInvertedLShape(position Vec2) Shape {
	return NewShape([]Vec2{
		//     [#]
		//      #
		//  # # #
		NewVec2(position.x+2, position.y+2),
		//      #
		//     [#]
		//  # # #
		NewVec2(position.x+2, position.y+1),
		//      #
		//      #
		// [#]# #
		NewVec2(position.x, position.y),
		//      #
		//      #
		//  #[#]#
		NewVec2(position.x+1, position.y),
		//      #
		//      #
		//  # #[#]
		NewVec2(position.x+2, position.y),
	})
}

func NewVerticalLineShape(position Vec2) Shape {
	return NewShape([]Vec2{
		// [#]
		//  #
		//  #
		//  #
		NewVec2(position.x, position.y+3),
		//  #
		// [#]
		//  #
		//  #
		NewVec2(position.x, position.y+2),
		//  #
		//  #
		// [#]
		//  #
		NewVec2(position.x, position.y+1),
		//  #
		//  #
		//  #
		// [#]
		NewVec2(position.x, position.y),
	})
}

func NewBoxShape(position Vec2) Shape {
	return NewShape([]Vec2{
		// [#]#
		//  # #
		NewVec2(position.x, position.y+1),
		// #[#]
		// # #
		NewVec2(position.x+1, position.y+1),
		//  # #
		// [#]#
		NewVec2(position.x, position.y),
		//  # #
		//  #[#]
		NewVec2(position.x+1, position.y),
	})
}
