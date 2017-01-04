package math

func PCount(n, k int) int {
	result := 1
	for i := n - k; i <= n; i++ {
		result *= i
	}
	return result
}

func extend(lhs, rhs []int) []int {
	result := make([]int, len(lhs)+len(rhs), len(lhs)+len(rhs))

	for i, _ := range lhs {
		result[i] = lhs[i]
	}

	for i, _ := range rhs {
		result[i+len(lhs)] = rhs[i]
	}

	return result
}

func factorial(n int) int64 {
	if n == 0 {
		return 1
	}

	var result int64

	result = int64(1)

	for i := int64(2); i <= int64(n); i += 1 {
		result *= i
	}
	return result
}

func P(lis []int, k int) [][]int {
	var result [][]int

	switch {
	// Stop case
	case k == 0:
		return make([][]int, 0, 0)

	// Trivial case
	case k == 1:
		result = make([][]int, len(lis), len(lis))

		for i, v := range lis {
			u := []int{v}
			result[i] = make([]int, 1, 1)
			copy(result[i], u)
		}

		return result
	}

	n := len(lis)
	result = make([][]int, 0, PCount(n, k))

	for i, v := range lis {
		current := []int{v}
		others := extend(lis[:i], lis[i+1:])

		for _, v := range P(others, k-1) {
			item := extend(current, v)
			result = append(result, item)
		}
	}

	return result
}
