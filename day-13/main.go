package main

import (
	"bufio"
	"os"
	"sort"
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

func findDecoderKeyPositions(nodes []Node) (int, int) {
	firstDividerPacket := ArrayNode{
		items: []Node{ArrayNode{
			items: []Node{NumberNode{value: 2}},
		}},
	}
	secondDividerPacket := ArrayNode{
		items: []Node{ArrayNode{
			items: []Node{NumberNode{value: 6}},
		}},
	}
	firstIndex := findInsertIndex(nodes, firstDividerPacket)
	secondIndex := findInsertIndex(nodes, secondDividerPacket)
	return firstIndex, secondIndex + 1
}

func findInsertIndex(nodes []Node, nodeToInsert Node) int {
	for index, node := range nodes {
		equality := nodeToInsert.Compare(node)
		if equality == IsLower {
			return index
		}
	}
	return len(nodes)
}

func getAllNodes(nodePairs []NodePair) []Node {
	nodes := make([]Node, 0)
	for _, pair := range nodePairs {
		nodes = append(nodes, pair.left, pair.right)
	}
	return nodes
}

func getDecoderKey(nodePairs []NodePair) int {
	nodes := getAllNodes(nodePairs)
	sortedNodes := sortNodes(nodes)
	first, second := findDecoderKeyPositions(sortedNodes)
	return (first + 1) * (second + 1)
}

func getRightOrderCount(nodePairs []NodePair) int {
	count := 0
	for index, pair := range nodePairs {
		equality := pair.left.Compare(pair.right)
		if equality == IsLower {
			count += index + 1
		}
	}
	return count
}

func sortNodes(nodes []Node) []Node {
	sort.SliceStable(nodes, func(i, j int) bool {
		left := nodes[i]
		right := nodes[j]
		return left.Compare(right) == IsLower
	})
	return nodes
}

func main() {
	lines := readLines()
	nodePairs := parseNodePairs(lines)
	println("Part 1 result:", getRightOrderCount(nodePairs))
	println("Part 2 result:", getDecoderKey(nodePairs))
}
