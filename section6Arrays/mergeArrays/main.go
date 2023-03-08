package merge_arrays

func mergeArrays(slc1, slc2 []int) []int {
	var mergeSlice []int

	for _, v := range slc1 {
		mergeSlice = append(mergeSlice, v)
	}

	for _, v := range slc2 {
		mergeSlice = append(mergeSlice, v)
	}

	for i := 0; i < len(mergeSlice); i++ {
		if i == len(mergeSlice)-1 {
			break
		}

		if mergeSlice[i] > mergeSlice[i+1] {
			mergeSlice[i], mergeSlice[i+1] = mergeSlice[i+1], mergeSlice[i]
		} else {
			continue
		}
	}

	return mergeSlice
}
