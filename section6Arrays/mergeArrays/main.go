package mergeArrays

import "sort"

func mergeArrays(slc1, slc2 []int) []int {
	var mergeSlice []int

	for _, v := range slc1 {
		mergeSlice = append(mergeSlice, v)
	}

	for _, v := range slc2 {
		mergeSlice = append(mergeSlice, v)
	}

	sort.Ints(mergeSlice)
	return mergeSlice
}
