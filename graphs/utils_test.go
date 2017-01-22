package math

import (
	"fmt"
	"testing"
)

func _() {
	fmt.Println("...")
}

func TestPick(t *testing.T) {
	m := map[string]float64{
		"a": 10.0,
		"b": 23.5,
		"c": 87.3,
		"d": 10.0,
	}

	if len(pick(&m, "a", "d")) != 2 {
		t.Fail()
	}

	if len(pick(&m, "whatever")) != 0 {
		t.Fail()
	}
}
