package piscine

func DescendAppendRange(max, min int) []int {
	arrayvalues := []int{}
	if max > min {
		for i := min + 1; i <= max; i++ {
			arrayvalues = append(arrayvalues, i)
		}
	} else {
		return arrayvalues
	}
	return arrayvalues
}
