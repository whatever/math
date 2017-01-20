package math

import (
	"math"
)

type WeightedUndiGraphEdge struct {
	weight float64
	src    *WeightedUndiGraphVertex
	dst    *WeightedUndiGraphVertex
}

type WeightedUndiGraphVertex struct {
	Label string
	Edges []*WeightedUndiGraphEdge
}

func (s *WeightedUndiGraphVertex) AddEdge(n *WeightedUndiGraphVertex, weight float64) {
	e := WeightedUndiGraphEdge{weight, s, n}
	s.Edges = append(s.Edges, &e)
}

/*
	Graph Helper
*/

type WeightedUGraph struct {
	vertices map[string]*WeightedUndiGraphVertex
}

func NewWeightedUGraph() WeightedUGraph {
	return WeightedUGraph{make(map[string]*WeightedUndiGraphVertex)}
}

func (s *WeightedUGraph) AddVertex(label string) *WeightedUGraph {
	n := WeightedUndiGraphVertex{label, nil}
	s.vertices[label] = &n
	return s
}

func (s *WeightedUGraph) AddEdge(src, dest string, weight float64) *WeightedUGraph {
	source := s.vertices[src]
	destination := s.vertices[dest]

	source.AddEdge(destination, weight)
	destination.AddEdge(source, weight)

	return s
}

func (s *WeightedUGraph) ShortestPath(src, dest string) []string {
	distances := make(map[string]float64)

	// Set "longest path" for all nodes
	for label, _ := range s.vertices {
		distances[label] = math.Inf(1)
	}
	distances[src] = 0

	return []string{}
}
