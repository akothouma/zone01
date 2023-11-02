package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var num, number, remainder int
	var runenumber rune
	i := 0
	for n > i {
		remainder = n % 10
		number += remainder
		num = n / 10
		n = num

		runenumber = rune(number)
		z01.PrintRune(runenumber)
	}
}
