package main

import (
	"strconv"
	"strings"
)

type Vec3 struct {
	x int
	y int
	z int
}

func NewVec3(x, y, z int) Vec3 {
	return Vec3{x, y, z}
}

func AddVec3(value *Vec3, x, y, z int) Vec3 {
	return NewVec3(value.x+x, value.y+y, value.z+z)
}

func parseInt(input string) int {
	value, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic("Invalid number")
	}
	return int(value)
}

func parseVec3(line string) Vec3 {
	parts := strings.Split(line, ",")
	return NewVec3(parseInt(parts[0]), parseInt(parts[1]), parseInt(parts[2]))
}
