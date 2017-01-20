package math

import (
	"fmt"
)

func extend(lhs, rhs []string) []string {
	result := make([]string, len(lhs)+len(rhs), len(lhs)+len(rhs))

	for i, _ := range lhs {
		result[i] = lhs[i]
	}

	for i, _ := range rhs {
		result[i+len(lhs)] = rhs[i]
	}

	return result
}

func IndexOfString(slice []string, needle string) int {
	for i, v := range slice {
		if v == needle {
			return i
		}
	}
	return -1
}

type DiGraphNode struct {
	Label    string
	Children []*DiGraphNode
}

func (self *DiGraphNode) ConnectNode(g *DiGraphNode) {
	for _, child := range self.Children {
		if child == g {
			return
		}
	}
	self.Children = append(self.Children, g)
}

func (self *DiGraphNode) GetEdgeLabels() []string {
	edgeLabels := make([]string, len(self.Children))
	for i, h := range self.Children {
		edgeLabels[i] = h.Label
	}
	return edgeLabels
}

type DiGraph struct {
	vertices map[string]*DiGraphNode
}

func NewDiGraph() DiGraph {
	return DiGraph{make(map[string]*DiGraphNode)}
}

func (self *DiGraph) CountNodes() int {
	return len(self.vertices)
}

func (self *DiGraph) CountEdges() int {
	edges := make(map[string][]string)

	for _, g := range self.vertices {
		for _, h := range g.Children {
			edges[g.Label] = append(edges[g.Label], h.Label)
		}
	}

	size := 0

	for _, v := range edges {
		size += len(v)
	}

	return size
}

func (self *DiGraph) AddNode(label string) {
	if _, ok := self.vertices[label]; ok {
		panic("A node already exists with that label")
	} else {
		self.vertices[label] = &DiGraphNode{label, nil}
	}
}

func (self *DiGraph) AddConnection(src, dest string) {
	g, g_ok := self.vertices[src]
	h, h_ok := self.vertices[dest]

	if g_ok && h_ok {
		g.ConnectNode(h)
	} else {
		panic("One of the nodes is missing for this edge")
	}
}

func (self *DiGraph) GetCyclesFor(label string, size int) [][]string {
	return getCycles(self.vertices[label], []string{}, size)
}

func getCycles(g *DiGraphNode, visited []string, size int) [][]string {

	switch {
	case g == nil || visited == nil:
		return [][]string{}

	case len(visited) > 0 && visited[0] == g.Label:
		newVisited := extend(visited, []string{g.Label})
		return [][]string{newVisited}

	case size == 0:
		return [][]string{}

	case IndexOfString(visited, g.Label) != -1:
		return [][]string{}
	}

	results := make([][]string, 0)
	newVisited := append(visited, g.Label)

	for _, h := range g.Children {

		newSize := size
		if newSize != -1 {
			newSize -= 1
		}

		for _, cycle := range getCycles(h, newVisited, newSize) {
			results = append(results, cycle)
		}
	}

	return results
}

func _() {
	fmt.Println("...")
}
