package main

func main() {
	input := ReadInput()
	firstField := ParseMapSquares(&input.rawMap)
	endField, facing := runMovement(firstField, input.movements)
	result := ((endField.coords.y + 1) * 1000) + ((endField.coords.x + 1) * 4) + int(facing)
	println("Day 1 result:", result)
}
