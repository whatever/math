package math

import (
	"testing"
)

func mapEquals(lhs, rhs map[int]int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for k, _ := range lhs {
		if lhs[k] != rhs[k] {
			return false
		}
	}

	return true
}

func TestMapEquals(t *testing.T) {
	a := map[int]int{
		1: 10,
		2: 2,
	}

	b := map[int]int{
		2: 2,
		1: 10,
	}

	c := map[int]int{
		2: 2,
		1: 10,
		3: 5,
	}

	if !mapEquals(a, b) {
		t.Fail()
	}

	if mapEquals(a, c) {
		t.Fail()
	}
}

func TestNaiveSieve(t *testing.T) {
	sieve := NewNaiveSieve(10)

	if len(sieve.Primes()) != 4 {
		t.Fail()
	}

	if !sieve.Contains(3) {
		t.Fail()
	}

	if sieve.Contains(6) {
		t.Fail()
	}
}

func TestFactorize(t *testing.T) {
	sieve := NewNaiveSieve(1000)
	if len(sieve.PrimeFactorize(13)) != 1 {
		t.Fail()
	}

	if !mapEquals(sieve.PrimeFactorize(13), map[int]int{13: 1}) {
		t.Fail()
	}

	if !mapEquals(sieve.PrimeFactorize(169), map[int]int{13: 2}) {
		t.Fail()
	}
}
