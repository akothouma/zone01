package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var num, number, remainder int
	var runenumber rune
	i := 0
	for n > i {
		num = n / 10
		remainder = n % 10
		number += remainder
		n = num

		runenumber = rune(number)
		z01.PrintRune(runenumber)
	}
	z01.PrintRune('0')
}
