package piscine

func IterativeFactorial(nb int) int {
	result := 0

	for i := nb; i > 0; nb-- {
		if nb == 0 {
			result := 1
			return result
		} else if nb > 63 || nb < 0 {
			return 0
		} else {
			result *= nb
		}
	}
	return result
}
