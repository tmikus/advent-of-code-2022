package main

import (
	"bufio"
	"os"
)

type Equality int
type NodeType int

const (
	IsLower   Equality = 1
	IsEqual            = 2
	IsGreater          = 3
)

const (
	Number NodeType = 1
	Array           = 2
)

type Node interface {
	Compare(other Node) Equality
	GetType() NodeType
	ToArrayNode() ArrayNode
}

type ArrayNode struct {
	items []Node
}

func (el ArrayNode) Compare(other Node) Equality {
	otherArray := other.ToArrayNode()
	for index := 0; index < len(el.items) && index < len(otherArray.items); index++ {
		equality := el.items[index].Compare(otherArray.items[index])
		if equality != IsEqual {
			return equality
		}
	}
	if len(el.items) < len(otherArray.items) {
		return IsLower
	}
	if len(el.items) == len(otherArray.items) {
		return IsEqual
	}
	return IsGreater
}

func (el ArrayNode) GetType() NodeType {
	return Array
}

func (el ArrayNode) ToArrayNode() ArrayNode {
	return el
}

type NumberNode struct {
	value int
}

func (el NumberNode) Compare(other Node) Equality {
	if other.GetType() == Number {
		otherNumber := other.(NumberNode)
		if el.value < otherNumber.value {
			return IsLower
		}
		if el.value == otherNumber.value {
			return IsEqual
		}
		return IsGreater
	}
	return el.ToArrayNode().Compare(other)
}

func (el NumberNode) GetType() NodeType {
	return Number
}

func (el NumberNode) ToArrayNode() ArrayNode {
	return ArrayNode{
		items: []Node{
			NumberNode{
				value: el.value,
			},
		},
	}
}

type NodePair struct {
	left  Node
	right Node
}

func parseNodePairs(lines []string) []NodePair {
	NodePairs := make([]NodePair, 0)
	for index := 0; index < len(lines); index += 3 {
		leftNodeScanner := NewLineScanner(lines[index])
		rightNodeScanner := NewLineScanner(lines[index+1])
		leftNode := ParseNode(&leftNodeScanner)
		rightNode := ParseNode(&rightNodeScanner)
		NodePairs = append(NodePairs, NodePair{
			left:  leftNode,
			right: rightNode,
		})
	}
	return NodePairs
}

func readLines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getRightOrderCount(NodePairs []NodePair) int {
	count := 0
	for index, pair := range NodePairs {
		equality := pair.left.Compare(pair.right)
		if equality == IsLower {
			count += (index + 1)
		}
	}
	return count
}

func main() {
	lines := readLines()
	nodePairs := parseNodePairs(lines)
	//for _, pair := range nodePairs {
	//	fmt.Printf("%v\n", pair.left)
	//	fmt.Printf("%v\n", pair.right)
	//	println("==================")
	//}
	println("Part 1 result:", getRightOrderCount(nodePairs))
}
