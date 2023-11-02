package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var num, remainder int

	i := 0
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for n > i {
			remainder = n % 10
			num = n / 10
			n = num
			z01.PrintRune(rune(remainder))
		}
	}
}
