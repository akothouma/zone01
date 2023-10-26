package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	a := 0
	b := a + 1
	c := b + 1

	for ; a <= 7; a++ {
		for ; b <= 8; b++ {
			for ; c <= 9; c++ {
				if a != 7 || b != 8 || c != 9 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
