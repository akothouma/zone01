package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := nb; i > 0; i-- {
		if i == 0 || i == 1 {
			return result
		} else if i > 63 || i < 0 {
			return 0
		} else {
			result *= i
		}
	}
	return result
}
