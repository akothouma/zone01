package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb == 0 || nb == 1 {
		return 1
	} else if nb > 63 || nb < 0 {
		return 0
	} else {
		result *= nb
	}
	return result
}
