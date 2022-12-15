package main

import (
	"strconv"
	"unicode"
)

type LineScanner struct {
	index int
	line  string
}

func NewLineScanner(line string) LineScanner {
	return LineScanner{
		index: 0,
		line:  line,
	}
}

func (s *LineScanner) GetChar() rune {
	result := s.PeekChar()
	s.index++
	return result
}

func (s *LineScanner) IsAtEnd() bool {
	return s.index == len(s.line)
}

func (s *LineScanner) PeekChar() rune {
	return rune(s.line[s.index])
}

func ParseNode(s *LineScanner) Node {
	if s.IsAtEnd() {
		panic("Reached the end of string before parsing a token")
	}
	c := s.PeekChar()
	if unicode.IsDigit(c) {
		return parseNumber(s)
	}
	if c != '[' {
		panic("Invalid character!")
	}
	s.index++ // Skip the [
	return parseArray(s)
}

func parseArray(s *LineScanner) Node {
	items := make([]Node, 0)
	for {
		if s.IsAtEnd() {
			panic("Reached end of string before end of array")
		}
		c := s.PeekChar()
		if c == ']' {
			s.index++ // Skip end of array
			break
		}
		if c == ',' {
			s.index++ // Skip commas
			continue
		}
		items = append(items, ParseNode(s))
	}
	return ArrayNode{items: items}
}

func parseNumber(s *LineScanner) Node {
	digits := ""
	for {
		if s.IsAtEnd() {
			break
		}
		digit := s.PeekChar()
		if !unicode.IsDigit(digit) {
			break
		}
		s.index++
		digits += string(digit)
	}
	number, err := strconv.ParseInt(digits, 10, 32)
	if err != nil {
		panic("Invalid number")
	}
	return NumberNode{value: int(number)}
}
