package piscine

func MakeRange(min, max int) []int {
	var requiredarray []int
	if min < max {
		size := max - min
		requiredarray = make([]int, size)

		for i := 0; i < size; i++ {
			if i == 0 {
				requiredarray[0] = min
			}
			min = min + 1
			requiredarray[i] = min
		}

	}
	return requiredarray
}
