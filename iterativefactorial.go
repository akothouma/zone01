package piscine

func IterativeFactorial(nb int) int {
	if nb == -1 || nb <= 0 {
		return 0
	} else if nb == 1 {
		return 1
	} else {
		return nb * IterativeFactorial(nb-1)
	}
}
