package piscine

func IterativeFactorial(nb int) int {
	result := 0

	for i := 0; i <= nb; i++ {
		if nb == 0 {
			return 1
		} else if nb > 63 || nb < 0 {
			return 0
		} else {
			result = nb * IterativeFactorial(nb-1)
		}
	}
	return result
}
