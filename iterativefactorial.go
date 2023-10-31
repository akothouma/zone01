package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := 0; i <= nb; i++ {
		if i == 0 || i == 1 {
			return 1
		} else if i > 63 || i < 0 {
			return 0
		} else {
			result = result * i
		}
	}
	return result
}
