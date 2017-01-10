package math

import (
	"math"
	"sort"
)

type NaiveSieve struct {
	sieve  map[int]bool
	primes []int
}

func NewNaiveSieve(limit int) NaiveSieve {

	/** Compute Prime Sieve **/
	sieve := make(map[int]bool)
	primes := make([]int, 1)

	sieve[2] = true
	primes[0] = 2

	for n := 3; n < limit; n += 2 {
		if IsNaivePrime(n) {
			sieve[n] = true
			primes = append(primes, n)
		}
	}

	return NaiveSieve{sieve, primes}
}

func (self *NaiveSieve) Contains(n int) bool {
	_, ok := self.sieve[n]
	return ok
}

func (self *NaiveSieve) PrimeFactorize(n int) map[int]int {
	factors := make(map[int]int)

	i := 0
	v := self.primes[i]

	for v <= n {

		if n%v == 0 {
			n /= v
			factors[v] += 1
		} else {
			// Update conditions
			i += 1
			v = self.primes[i]
		}
	}

	return factors
}

func (self *NaiveSieve) Primes() []int {
	primes := make([]int, 0, len(self.sieve))
	for n, _ := range self.sieve {
		primes = append(primes, n)
	}
	sort.Ints(primes)
	return primes
}

func IsNaivePrime(n int) bool {
	switch {
	case n == 1:
		return false
	case n == 2:
		return true
	case n%2 == 0:
		return false
	}

	limit := int(math.Sqrt(float64(n)))

	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
