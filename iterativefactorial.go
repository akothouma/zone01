package piscine

func IterativeFactorial(nb int) int {
	if nb == 1 {
		return 1
	} else if nb > 1 {
		return nb * IterativeFactorial(nb-1)
	} else  if nb > 63 || nb < - 63{
		return 0
	}
}



}
