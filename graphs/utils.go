package math

import (
	"math"
)

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

func unpick(values *map[string]float64, keys ...string) map[string]float64 {
	picked := make(map[string]float64)

	for label, v := range *values {
		picked[label] = v
	}

	keyMap := make(map[string]bool)

	for _, label := range keys {
		keyMap[label] = true
	}

	for label, _ := range *values {
		_, exists := keyMap[label]

		if exists {
			delete(picked, label)
		}
	}

	return picked
}

func pick(values *map[string]float64, keys ...string) map[string]float64 {
	picked := make(map[string]float64)

	for _, v := range keys {
		value, exists := (*values)[v]
		if exists {
			picked[v] = value
		}
	}

	return picked
}

func keys(values *map[string]float64) []string {
	results := make([]string, len(*values))
	i := 0

	for label, _ := range *values {
		results[i] = label
		i++
	}

	return results
}

func keysBool(values *map[string]bool) []string {
	results := make([]string, len(*values))
	i := 0

	for label, _ := range *values {
		results[i] = label
		i++
	}

	return results
}
