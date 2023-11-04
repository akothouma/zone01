package piscine

func AppendRange(min, max int) []int {
	var requiredarray []int
	if min <= max {
		for i := min; i < max; i++ {
			requiredarray = append(requiredarray, i)
		}
	}
	return requiredarray
}
