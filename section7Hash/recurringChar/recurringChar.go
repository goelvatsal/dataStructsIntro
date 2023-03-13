package recurringchar

type recurringCharS1 struct{}

func (s1 recurringCharS1) findRecurringChar(ints []int) int {
	charSet := make(map[int]bool, 0)
	result := -1

	for _, v := range ints {
		if _, ok := charSet[v]; !ok {
			charSet[v] = true
		} else {
			result = v
			break
		}
	}

	return result
}
