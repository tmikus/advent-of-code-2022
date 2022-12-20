package main

func tagCubesFacingOutside(size Vec3, cubeMap *[][][]*Cube) {
	tagXSides(size, cubeMap, 0)
	tagXSides(size, cubeMap, size.x-1)

	tagYSides(size, cubeMap, 0)
	tagYSides(size, cubeMap, size.y-1)

	tagZSides(size, cubeMap, 0)
	tagZSides(size, cubeMap, size.z-1)
}

func tagXSides(size Vec3, cubeMap *[][][]*Cube, x int) {
	for y := 0; y < size.y; y++ {
		for z := 0; z < size.z; z++ {
			recursivelyTagNeighbours((*cubeMap)[x][y][z])
		}
	}
}

func tagYSides(size Vec3, cubeMap *[][][]*Cube, y int) {
	for x := 0; x < size.x; x++ {
		for z := 0; z < size.z; z++ {
			recursivelyTagNeighbours((*cubeMap)[x][y][z])
		}
	}
}

func tagZSides(size Vec3, cubeMap *[][][]*Cube, z int) {
	for x := 0; x < size.x; x++ {
		for y := 0; y < size.y; y++ {
			recursivelyTagNeighbours((*cubeMap)[x][y][z])
		}
	}
}

func recursivelyTagNeighbours(cube *Cube) {
	if cube == nil || !cube.void || cube.visited {
		return
	}
	cube.isFacingOutside = true
	cube.visited = true
	recursivelyTagNeighbours(cube.left)
	recursivelyTagNeighbours(cube.right)
	recursivelyTagNeighbours(cube.front)
	recursivelyTagNeighbours(cube.back)
	recursivelyTagNeighbours(cube.top)
	recursivelyTagNeighbours(cube.bottom)
}
