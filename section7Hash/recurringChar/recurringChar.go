package recurringchar

type recurringCharS1 struct{}

func (s1 recurringCharS1) findRecurringChar(input1 []int) int {
	kvPair := make(map[int]bool, 0)
	var result int

	for _, v := range input1 {
		result = v
		if _, ok := kvPair[v]; ok != true {
			kvPair[v] = false
		} else {
			kvPair[v] = true
			break
		}
	}
	return result
}
