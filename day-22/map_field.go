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
	longestRow := 0
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		row := parseMapFields(line)
		if len(row) > longestRow {
			longestRow = len(row)
		}
		mapFields = append(mapFields, row)
	}
	addPaddingToRows(&mapFields, longestRow)
	return mapFields
}

func addPaddingToRows(fields *[][]MapField, width int) {
	for rowIndex := 0; rowIndex < len(*fields); rowIndex++ {
		row := (*fields)[rowIndex]
		rowWidth := len(row)
		if rowWidth > width {
			panic("The row should not be longer than the max width!")
		}
		if rowWidth == width {
			continue
		}
		missingFields := width - rowWidth
		for x := 0; x < missingFields; x++ {
			row = append(row, None)
		}
		(*fields)[rowIndex] = row
	}
}
