package math

import (
	"fmt"
	"math"
)

func MinElement(values *map[string]float64) (string, float64) {
	minVal := math.Inf(1)
	minInd := ""

	for label, val := range *values {
		if val <= minVal {
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

func (s *WeightedUndiGraphEdge) String() string {
	return fmt.Sprintf("[%s %s]: %0.1f", s.src.Label, s.dst.Label, s.weight)
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
	unvisited := make(map[string]bool)

	// Set "longest path" for all nodes
	for label, _ := range s.vertices {
		distances[label] = math.Inf(1)
		previous[label] = math.Inf(1)
		unvisited[label] = true
	}
	distances[src] = 0.0

	// ...
	for len(distances) > 0 {

		// Find the next element to check, and remove it
		currentLabel, currentValue := MinElement(&distances)
		current := s.vertices[currentLabel]

		delete(distances, currentLabel)

		// ...
		fmt.Println("-->")
		for _, e := range current.Edges {
			fmt.Println(e)
		}

		_, _, _ = currentLabel, currentValue, current

		// fmt.Println(current, currentValue, current.Edges)
	}

	return []string{}
}

func _() {
	fmt.Println()
}
