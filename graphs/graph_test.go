package math

import (
	"fmt"
	"strings"
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
	g.AddNode("b")
	g.AddNode("c")
	g.AddNode("d1")
	g.AddNode("d2")
	g.AddConnection("a", "b")
	g.AddConnection("b", "c")
	g.AddConnection("c", "d1")
	g.AddConnection("d1", "a")
	g.AddConnection("c", "d2")
	g.AddConnection("d2", "a")
	// g.AddConnection("d", "b")

	// g.AddConnection("e", "a")
	// g.AddConnection("a", "c")
	// g.AddConnection("e", "f")
	// g.AddConnection("a", "B2")
	// g.AddConnection("B2", "c")
	// g.AddConnection("B2", "d")
	// g.AddConnection("c", "a")

	cycles3 := getCycles(g.vertices["a"], []string{"a", "b", "c", "d2"}, -1)

	if strings.Join(cycles3[0], ",") != "a,b,c,d2,a" {
		t.Fail()
	}

	cycles := getCycles(g.vertices["a"], []string{}, -1)

	if len(cycles) != 2 {
		t.Fail()
	}

	cycles2 := getCycles(nil, nil, -1)

	if len(cycles2) != 0 {
		t.Fail()
	}

	g2 := NewGraph()
	g2.AddNode("a")
	g2.AddNode("b")
	g2.AddNode("c")
	g2.AddNode("d")
	g2.AddNode("f")
	g2.AddNode("g")
	g2.AddConnection("a", "b")
	g2.AddConnection("b", "c")
	g2.AddConnection("c", "a")
	g2.AddConnection("c", "d")
	g2.AddConnection("d", "a")
	g2.AddConnection("a", "f")
	g2.AddConnection("f", "g")
	g2.AddConnection("g", "a")

	cycles4 := g2.GetCyclesFor("a", -1)

	if len(cycles4) != 3 {
		t.Fail()
	}

	g3 := NewGraph()
	g3.AddNode("a")
	g3.AddNode("b")
	g3.AddNode("c")
	g3.AddConnection("b", "c")
	g3.AddConnection("c", "b")

	noCycles := g3.GetCyclesFor("a", -1)

	if len(noCycles) != 0 {
		t.Fail()
	}

	oneCycle := g3.GetCyclesFor("b", -1)

	if len(oneCycle) != 1 {
		t.Fail()
	}
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
