package math

import (
	"testing"
)

func TestExtend(t *testing.T) {
	a := extend([]int{}, []int{})

	if len(a) != 0 {
		t.Fail()
	}

	b := extend([]int{1, 2, 3}, []int{4, 5})

	if len(b) != 5 {
		t.Fail()
	}

	if b[3] != 4 {
		t.Fail()
	}

	if b[4] != 5 {
		t.Fail()
	}
}

func TestX(t *testing.T) {
}

func IsEqualIntArray(lhs, rhs []int) bool {
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

func IsEqualIntArrayArray(lhs, rhs [][]int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i, _ := range lhs {
		if !IsEqualIntArray(lhs[i], rhs[i]) {
			return false
		}
	}

	return true
}

func TestC(t *testing.T) {
	vals := []int{1, 2, 3}

	if !IsEqualIntArrayArray(P(vals, 0), [][]int{}) {
		t.Fail()
	}

	x := [][]int{[]int{1}, []int{2}, []int{3}}

	if !IsEqualIntArrayArray(P(vals, 1), x) {
		t.Fail()
	}

	y := P(vals, 1)

	if len(y) != 3 {
		t.Fail()
	}

	z := P(vals, 2)

	if len(z) != 6 {
		t.Fail()
	}

	vals2 := []int{0, 1, 2, 3, 4}
	a := P(vals2, 5)

	if len(a) != 120 {
		t.Fail()
	}
}
