package math

import (
	"fmt"
	"math"
	"testing"
)

func _() {
	fmt.Println("...")
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

func TestMinElement(t *testing.T) {
	a := map[string]float64{
		"red":  23.0,
		"blue": -200.0,
		"y":    200.0,
	}
	label, result := MinElement(&a)
	if label != "blue" || result != -200.0 {
		t.Fail()
	}
}

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

func TestPick(t *testing.T) {
	m := map[string]float64{
		"a": 10.0,
		"b": 23.5,
		"c": 87.3,
		"d": 10.0,
	}

	keys := []string{"b", "c", "d"}

	if len(pick(&m, "a", "d")) != 2 {
		t.Fail()
	}

	if len(pick(&m, "whatever")) != 0 {
		t.Fail()
	}

	picked := pick(&m, keys...)

	if len(picked) != 3 {
		t.Fail()
	}

	if picked["b"] != 23.5 {
		t.Fail()
	}
}
