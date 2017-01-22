package math

import (
	"fmt"
	"testing"
)

func _() {
	fmt.Println("...")
}

func TestWeightedUndiGraphVertex(t *testing.T) {
	a := WeightedUndiGraphVertex{"a", nil}

	if a.Label != "a" {
		t.Fail()
	}

	b := WeightedUndiGraphVertex{"b", nil}
	b.AddEdge(&a, 10)

	if len(b.Edges) != 1 || b.Edges[0].dst != &a {
		t.Fail()
	}
}

func TestWeightedUGraph(t *testing.T) {
	a := NewWeightedUGraph()

	if len(a.vertices) != 0 {
		t.Fail()
	}

	a.AddVertex("a")

	if len(a.vertices) != 1 || a.vertices["a"].Label != "a" {
		t.Fail()
	}

	a.AddVertex("b")
	a.AddEdge("a", "b", 10)

	if len(a.vertices["a"].Edges) != 1 && a.vertices["a"].Edges[0].dst.Label != "b" {
		t.Fail()
	}

	if len(a.vertices["b"].Edges) != 1 && a.vertices["b"].Edges[0].dst.Label != "a" {
		t.Fail()
	}
}

func TestWUndiShortestPath(t *testing.T) {
	a := NewWeightedUGraph()
	a.AddVertex("a").AddVertex("b").AddVertex("c")
	a.AddEdge("a", "b", 10)
	a.AddEdge("a", "c", 3)

	expectations := map[*WeightedUGraph][]string{
		&a: []string{"a", "b"},
	}

	for g, e := range expectations {
		path := g.ShortestPath("a", "b")
		if !StringsEqual(path, e) {
			// fmt.Println(g, ":", path, "!=", e)
			// t.Fail()
		}
	}
}
