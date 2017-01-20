package math

import (
	"fmt"
	"math"
)

func MinElement(values *map[string]float64) (string, float64) {
	minVal := math.Inf(1)
	minInd := ""

	for label, val := range *values {
		if val < minVal {
			minVal = val
			minInd = label
		}
	}
	return minInd, minVal
}

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
	previous := make(map[string]float64)
	visited := make(map[string]bool)

	// Set "longest path" for all nodes
	for label, _ := range s.vertices {
		distances[label] = math.Inf(1)
		previous[label] = math.Inf(1)
		visited[label] = false
	}
	distances[src] = 0

	// ...
	current := *s.vertices[src]

	// ...
	for _, e := range current.Edges {
		currentDist := distances[e.dst.Label]
		dist := math.Min(e.weight, currentDist)
		fmt.Println(current.Label, e.dst.Label, currentDist, dist)
	}

	return []string{}
}
