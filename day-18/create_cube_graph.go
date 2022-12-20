package main

func createCubeGraph(positions []Vec3) (Vec3, []Cube, [][][]*Cube) {
	mapSize := getCubeMapSize(&positions)
	cubes, cubeMap := createCubeMap(positions, mapSize)
	for index := 0; index < len(cubes); index++ {
		cube := &cubes[index]
		updateAdjacencyPointers(cube, &cubeMap, mapSize)
	}
	return mapSize, cubes, cubeMap
}

// X->Y->Z
func createCubeMap(positions []Vec3, size Vec3) ([]Cube, [][][]*Cube) {
	cubes := initCubes(size)
	cubeMap := initCubeMap(&cubes, size)
	for _, position := range positions {
		cube := cubeMap[position.x][position.y][position.z]
		cube.void = false
	}
	return cubes, cubeMap
}

func getCube(cubeMap *[][][]*Cube, mapSize Vec3, position Vec3) *Cube {
	if position.x < 0 ||
		position.x >= mapSize.x ||
		position.y < 0 ||
		position.y >= mapSize.y ||
		position.z < 0 ||
		position.z >= mapSize.z {
		return nil
	}
	return (*cubeMap)[position.x][position.y][position.z]
}

func getCubeMapSize(positions *[]Vec3) Vec3 {
	maxX := 0
	maxY := 0
	maxZ := 0
	for index := 0; index < len(*positions); index++ {
		position := &(*positions)[index]
		if position.x > maxX {
			maxX = position.x
		}
		if position.y > maxY {
			maxY = position.y
		}
		if position.z > maxZ {
			maxZ = position.z
		}
	}
	return NewVec3(maxX+1, maxY+1, maxZ+1)
}

func initCubeMap(cubes *[]Cube, size Vec3) [][][]*Cube {
	cubeMap := make([][][]*Cube, 0)
	for x := 0; x < size.x; x++ {
		inner := make([][]*Cube, 0)
		for y := 0; y < size.y; y++ {
			row := make([]*Cube, size.z)
			for z := 0; z < size.z; z++ {
				cube := &(*cubes)[(x*size.y*size.z)+(y*size.z)+z]
				row[z] = cube
			}
			inner = append(inner, row)
		}
		cubeMap = append(cubeMap, inner)
	}
	return cubeMap
}

func initCubes(size Vec3) []Cube {
	cubes := make([]Cube, 0)
	for x := 0; x < size.x; x++ {
		for y := 0; y < size.y; y++ {
			for z := 0; z < size.z; z++ {
				cubes = append(cubes, NewCube(NewVec3(x, y, z)))
			}
		}
	}
	return cubes
}

func updateAdjacencyPointers(cube *Cube, cubeMap *[][][]*Cube, mapSize Vec3) {
	cube.front = getCube(cubeMap, mapSize, AddVec3(&cube.position, 0, 0, -1))
	if cube.front != nil {
		cube.front.back = cube
	}
	cube.back = getCube(cubeMap, mapSize, AddVec3(&cube.position, 0, 0, 1))
	if cube.back != nil {
		cube.back.front = cube
	}
	cube.top = getCube(cubeMap, mapSize, AddVec3(&cube.position, 0, 1, 0))
	if cube.top != nil {
		cube.top.bottom = cube
	}
	cube.bottom = getCube(cubeMap, mapSize, AddVec3(&cube.position, 0, -1, 0))
	if cube.bottom != nil {
		cube.bottom.top = cube
	}
	cube.left = getCube(cubeMap, mapSize, AddVec3(&cube.position, -1, 0, 0))
	if cube.left != nil {
		cube.left.right = cube
	}
	cube.right = getCube(cubeMap, mapSize, AddVec3(&cube.position, 1, 0, 0))
	if cube.right != nil {
		cube.right.left = cube
	}
}
