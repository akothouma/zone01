package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 63 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		for i := nb; i > 2; i-- {
			result = result * i
		}
	}
	return result
}
