package math

import (
	_ "fmt"
)

type GraphNode struct {
	Label    string
	Children []*GraphNode
}

func (self *GraphNode) ConnectNode(g *GraphNode) {
	self.Children = append(self.Children, g)
}

type Graph struct {
	vertices map[string]*GraphNode
}

func NewGraph() Graph {
	return Graph{make(map[string]*GraphNode)}
}

func (self *Graph) CountNodes() int {
	return len(self.vertices)
}

func (self *Graph) CountEdges() int {
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

func (self *Graph) AddNode(label string) {
	if _, ok := self.vertices[label]; ok {
		panic("A node already exists with that label")
	} else {
		self.vertices[label] = &GraphNode{label, nil}
	}
}

func (self *Graph) AddConnection(src, dest string) {
	g, g_ok := self.vertices[src]
	h, h_ok := self.vertices[dest]

	if g_ok && h_ok {
		g.ConnectNode(h)
	} else {
		panic("One of the nodes is missing for this edge")
	}
}

func (self *Graph) GetNaiveCycles() {
}
