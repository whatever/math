package math

import (
	"fmt"
	"testing"
)

func TestGraphNode(t *testing.T) {
	g := GraphNode{"g", nil}
	h := GraphNode{"h", nil}

	if len(g.Children) != 0 {
		t.Fail()
	}

	g.ConnectNode(&h)

	if len(g.Children) != 1 {
		t.Fail()
	}

	if g.Children[0].Label != "h" {
		t.Fail()
	}

	h.Label = "whatever"

	if g.Children[0].Label != "whatever" {
		t.Fail()
	}
}

func TestGraph(t *testing.T) {
	g := NewGraph()

	if g.CountNodes() != 0 {
		t.Fail()
	}

	g.AddNode("start")
	g.AddNode("end")

	if g.CountNodes() != 2 {
		t.Fail()
	}

	if g.CountEdges() != 0 {
		t.Fail()
	}

	g.AddConnection("start", "end")

	if g.CountEdges() != 1 {
		fmt.Println(g.CountEdges())
		t.Fail()
	}

	g.AddConnection("end", "start")

	if g.CountEdges() != 2 {
		fmt.Println(g.CountEdges())
		t.Fail()
	}

	g.AddNode("whatever")
	g.AddConnection("whatever", "end")
	g.AddConnection("whatever", "start")
	g.AddConnection("start", "whatever")

	if g.CountEdges() != 5 {
		fmt.Println(g.CountEdges())
		t.Fail()
	}
}

func TestSimpleGraph(t *testing.T) {
	g := NewGraph()
	g.AddNode("a")
	g.AddNode("B1")
	g.AddNode("B2")
	g.AddNode("c")
	g.AddConnection("a", "B1")
	g.AddConnection("a", "B2")
	g.AddConnection("B1", "c")

	cycles := getCycles(g.vertices["a"], []string{})

	fmt.Println(cycles)
}

func TestGraphMethods(t *testing.T) {
	g := NewGraph()
	g.AddNode("a")
	g.AddNode("b")
	g.AddNode("c")
	g.AddNode("d")
	g.AddConnection("a", "b")
	g.AddConnection("b", "c")
	g.AddConnection("c", "d")
	g.AddConnection("d", "a")

	if g.CountNodes() != 4 {
		t.Fail()
	}

	if g.CountEdges() != 4 {
		t.Fail()
	}
}

func _() {
	fmt.Println("NOOP")
}
