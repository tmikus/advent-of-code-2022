package main

import (
	"bufio"
	"os"
)

type Node struct {
	height           int
	index            int
	validNodeIndices []int
}

func newNode(height int, index int) Node {
	return Node{
		height:           height,
		index:            index,
		validNodeIndices: make([]int, 0),
	}
}

type NodeGraph struct {
	nodes           []Node
	startNodeIndex  int
	targetNodeIndex int
}

type NodeGrid struct {
	nodes           []Node
	startNodeIndex  int
	targetNodeIndex int
	width           int
}

func cloneArray(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)
	return result
}

func contains(s *[]int, e int) bool {
	for index := 0; index < len(*s); index++ {
		if (*s)[index] == e {
			return true
		}
	}
	return false
}

func findShortestPath(graph NodeGraph) []int {
	paths := [][]int{
		{graph.startNodeIndex},
	}
	visitedNodes := []int{graph.startNodeIndex}
	for {
		if len(paths) == 0 {
			panic("Couldn't find the path to the end!")
		}
		nextPaths := make([][]int, 0)
		for pathIndex := 0; pathIndex < len(paths); pathIndex++ {
			path := paths[pathIndex]
			lastNodeIndex := path[len(path)-1]
			lastNode := &graph.nodes[lastNodeIndex]
			for _, targetNodeIndex := range lastNode.validNodeIndices {
				if contains(&visitedNodes, targetNodeIndex) {
					continue
				}
				newPath := cloneArray(path)
				newPath = append(newPath, targetNodeIndex)
				visitedNodes = append(visitedNodes, targetNodeIndex)
				if graph.targetNodeIndex == targetNodeIndex {
					return newPath
				}
				nextPaths = append(nextPaths, newPath)
			}
		}
		paths = nextPaths
	}
}

func parseNodesGrid(lines []string) NodeGrid {
	nodes := make([]Node, 0)
	startNodeIndex := 0
	targetNodeIndex := 0
	width := len(lines[0])
	for _, line := range lines {
		for index := 0; index < width; index++ {
			nodeChar := line[index]
			nodeIndex := len(nodes)
			var node Node
			if nodeChar == 'S' {
				node = newNode(0, nodeIndex)
				startNodeIndex = nodeIndex
			} else if nodeChar == 'E' {
				node = newNode(int('z')-int('a'), nodeIndex)
				targetNodeIndex = nodeIndex
			} else {
				height := int(nodeChar) - int('a')
				node = newNode(height, nodeIndex)
			}
			nodes = append(nodes, node)
		}
	}
	return NodeGrid{
		nodes:           nodes,
		startNodeIndex:  startNodeIndex,
		targetNodeIndex: targetNodeIndex,
		width:           width,
	}
}

func parseNodeGraph(nodeGrid NodeGrid) NodeGraph {
	for index := 0; index < len(nodeGrid.nodes); index++ {
		updateNodeLinks(&nodeGrid.nodes[index], &nodeGrid)
	}
	return NodeGraph{
		nodes:           nodeGrid.nodes,
		startNodeIndex:  nodeGrid.startNodeIndex,
		targetNodeIndex: nodeGrid.targetNodeIndex,
	}
}

func readLines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func updateNodeWhenLinkExists(node *Node, nextNode *Node) {
	if nextNode.height-node.height > 1 {
		return
	}
	node.validNodeIndices = append(node.validNodeIndices, nextNode.index)
}

func updateNodeLinks(node *Node, nodeGrid *NodeGrid) {
	positionWithinRow := node.index % nodeGrid.width
	// Has left node
	if positionWithinRow > 0 {
		updateNodeWhenLinkExists(node, &nodeGrid.nodes[node.index-1])
	}
	// Has right node
	if positionWithinRow < nodeGrid.width-1 {
		updateNodeWhenLinkExists(node, &nodeGrid.nodes[node.index+1])
	}
	// Has up node
	if node.index >= nodeGrid.width {
		updateNodeWhenLinkExists(node, &nodeGrid.nodes[node.index-nodeGrid.width])
	}
	// Has down node
	if node.index < len(nodeGrid.nodes)-nodeGrid.width {
		updateNodeWhenLinkExists(node, &nodeGrid.nodes[node.index+nodeGrid.width])
	}
}

func main() {
	lines := readLines()
	grid := parseNodesGrid(lines)
	graph := parseNodeGraph(grid)
	path := findShortestPath(graph)
	println("Shortest path:", len(path)-1)
}
