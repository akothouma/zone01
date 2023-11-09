package piscine

func DescendAppendRange(max, min int) []int {
	arrayvalues := []int{}
	if max > min {
		for i := max; i > min; i++ {
			arrayvalues = append(arrayvalues, i)
		}
	} else {
		return arrayvalues
	}
	return arrayvalues
}
