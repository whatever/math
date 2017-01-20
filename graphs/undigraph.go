package math

type WeightedUndiGraphEdge struct {
	weight int
	src    *WeightedUndiGraphVertex
	dst    *WeightedUndiGraphVertex
}

type WeightedUndiGraphVertex struct {
	Label string
	Edges []*WeightedUndiGraphEdge
}

func (s *WeightedUndiGraphVertex) AddEdge(n *WeightedUndiGraphVertex, weight int) {
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

func (s *WeightedUGraph) AddEdge(src, dest string, weight int) *WeightedUGraph {
	source := s.vertices[src]
	destination := s.vertices[dest]

	source.AddEdge(destination, weight)
	destination.AddEdge(source, weight)

	return s
}

func (s *WeightedUGraph) ShortestPath(src, dest string) []string {
	return []string{}
}
