package mergeArrays

import "sort"

func mergeArrays(slc1, slc2 []int) []int {
	var mergeSlice []int

	mergeSlice = append(mergeSlice, slc1...)
	mergeSlice = append(mergeSlice, slc2...)

	sort.Ints(mergeSlice)
	return mergeSlice
}
