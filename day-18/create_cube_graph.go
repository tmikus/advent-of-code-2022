package main

func createCubeGraph(positions []Vec3) []Cube {
	cubes := createCubes(positions)
	mapSize := getCubeMapSize(&cubes)
	cubeMap := createCubeMap(&cubes, mapSize)
	for index := 0; index < len(cubes); index++ {
		cube := &cubes[index]
		updateAdjacencyPointers(cube, &cubeMap, mapSize)
	}
	return cubes
}

// X->Y->Z
func createCubeMap(cubes *[]Cube, mapSize Vec3) [][][]*Cube {
	cubeMap := initCubeMap(mapSize)
	for index := 0; index < len(*cubes); index++ {
		cube := &(*cubes)[index]
		cubeMap[cube.position.x][cube.position.y][cube.position.z] = cube
	}
	return cubeMap
}

func createCubes(positions []Vec3) []Cube {
	cubes := make([]Cube, 0)
	for _, position := range positions {
		cubes = append(cubes, NewCube(position))
	}
	return cubes
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

func getCubeMapSize(cubes *[]Cube) Vec3 {
	maxX := 0
	maxY := 0
	maxZ := 0
	for index := 0; index < len(*cubes); index++ {
		cube := &(*cubes)[index]
		if cube.position.x > maxX {
			maxX = cube.position.x
		}
		if cube.position.y > maxY {
			maxY = cube.position.y
		}
		if cube.position.z > maxZ {
			maxZ = cube.position.z
		}
	}
	return NewVec3(maxX+1, maxY+1, maxZ+1)
}

func initCubeMap(size Vec3) [][][]*Cube {
	result := make([][][]*Cube, 0)
	for x := 0; x <= size.x; x++ {
		inner := make([][]*Cube, 0)
		for y := 0; y <= size.y; y++ {
			inner = append(inner, make([]*Cube, size.z))
		}
		result = append(result, inner)
	}
	return result
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
