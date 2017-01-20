package math

type Graph interface {
	AddVertex(label string)
	AddEdge(src, destination string)
	CountVertices() int
	CountEdges() int
}
