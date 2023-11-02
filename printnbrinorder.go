package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var num, remainder int
	var runenumber rune
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for n > 0 {
			remainder = n % 10
			num = n / 10
			runenumber = rune(remainder)
			n = num
			z01.PrintRune(runenumber)
		}
	}
}
