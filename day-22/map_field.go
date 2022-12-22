package main

import "bufio"

type MapField int

const (
	None MapField = 0
	Path          = 1
	Wall          = 2
)

func parseMapFields(line string) []MapField {
	fields := make([]MapField, 0)
	for index := 0; index < len(line); index++ {
		field := None
		switch line[index] {
		case '.':
			field = Path
		case '#':
			field = Wall
		}
		fields = append(fields, field)
	}
	return fields
}

func readMapFields(scanner *bufio.Scanner) [][]MapField {
	mapFields := make([][]MapField, 0)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		mapFields = append(mapFields, parseMapFields(line))
	}
	return mapFields
}
