package math

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
