package math

import (
	"fmt"
	"testing"
)

func _() {
	fmt.Println("...")
}

func TestWeightedUndiGraphVertex(t *testing.T) {
	a := WeightedUndiGraphVertex{"a", 10, nil}

	if a.Weight != 10 {
		t.Fail()
	}

	b := WeightedUndiGraphVertex{"b", 10, nil}
	b.AddEdge(&a)

	if len(b.Neighbors) != 1 || b.Neighbors[0] != &a {
		t.Fail()
	}
}

func TestWeightedUGraph(t *testing.T) {
	a := NewWeightedUGraph()

	if len(a.vertices) != 0 {
		t.Fail()
	}

	a.AddVertex("a", 10)

	if len(a.vertices) != 1 || a.vertices["a"].Weight != 10 {
		t.Fail()
	}

	a.AddVertex("b", 4)
	a.AddEdge("a", "b")

	if len(a.vertices["a"].Neighbors) != 1 && a.vertices["a"].Neighbors[0].Label != "b" {
		t.Fail()
	}

	if len(a.vertices["b"].Neighbors) != 1 && a.vertices["b"].Neighbors[0].Label != "a" {
		t.Fail()
	}
}

func TestWUndiShortestPath(t *testing.T) {

	a := NewWeightedUGraph()
	a.AddVertex("a", 10).AddVertex("b", 23)

	expectations := map[*WeightedUGraph][]string{
		&a: []string{"a", "b"},
	}

	for g, e := range expectations {
		if !StringsEqual(g.ShortestPath("a", "b"), e) {
			fmt.Println(g, "!=", e)
			t.Fail()
		}
	}
}

func StringsEqual(lhs, rhs []string) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i, _ := range lhs {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}
