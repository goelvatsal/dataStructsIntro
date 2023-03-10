package mergearraysandsort

type mergeArraysS1 struct{}

func (s1 mergeArraysS1) mergeArrays(arr1, arr2 []int) []int {
	var merge []int

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		a1, a2 := arr1[i], arr2[j]
		if a1 == a2 {
			merge = append(merge, a1, a2)
			i++
			j++
		} else if a1 > a2 {
			merge = append(merge, a2)
			j++
		} else {
			merge = append(merge, a1)
			i++
		}
	}

	for ; i < len(arr1); i++ {
		merge = append(merge, arr1[i])
	}

	for ; j < len(arr2); j++ {
		merge = append(merge, arr2[j])
	}
	return merge
}
