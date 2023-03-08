package revstr

func revStr(str string) string {
	slc := []rune(str)
	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]
	}

	return string(slc)
}
