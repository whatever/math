package math

import (
	"fmt"
	"math"
)

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

/*
Djikstra's Shortest Path Algorithm (Poorly) Implemented
Return a map of distances
*/
func (s *WeightedUGraph) ShortestDistances(src string) map[string]float64 {
	visited := make(map[string]bool)

	// Create vertex set
	q := make(map[string]bool)
	distances := make(map[string]float64)
	prev := make(map[string]*WeightedUndiGraphVertex)

	for _, v := range s.vertices {
		q[v.Label] = true
		distances[v.Label] = math.Inf(1)
		prev[v.Label] = nil
	}

	distances[src] = 0

	for len(q) > 0 {
		unvisitedDistances := unpick(&distances, keysBool(&visited)...)
		label, _ := MinElement(&unvisitedDistances)
		delete(q, label)
		visited[label] = true

		u := s.vertices[label]

		for _, e := range u.Edges {
			_, ok := q[e.dst.Label]

			if !ok {
				continue
			}

			alt := distances[u.Label] + e.weight
			v := e.dst

			if alt < distances[v.Label] {
				distances[v.Label] = alt
				prev[v.Label] = u
			}
		}
	}

	return distances
}

func _() {
	fmt.Println()
}
