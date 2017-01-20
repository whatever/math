package math

type WeightedUndiGraphVertex struct {
	Label     string
	Weight    int
	Neighbors []*WeightedUndiGraphVertex
}

func (s *WeightedUndiGraphVertex) AddEdge(n *WeightedUndiGraphVertex) {
	s.Neighbors = append(s.Neighbors, n)
}

type WeightedUGraph struct {
	vertices map[string]*WeightedUndiGraphVertex
}

func NewWeightedUGraph() WeightedUGraph {
	return WeightedUGraph{make(map[string]*WeightedUndiGraphVertex)}
}

func (s *WeightedUGraph) AddVertex(label string, weight int) *WeightedUGraph {
	n := WeightedUndiGraphVertex{label, weight, nil}
	s.vertices[label] = &n
	return s
}

func (s *WeightedUGraph) AddEdge(src, dest string) *WeightedUGraph {
	source := s.vertices[src]
	destination := s.vertices[dest]

	source.AddEdge(destination)
	destination.AddEdge(source)
	return s
}
