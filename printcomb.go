package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	a := '0'
	b := a + 1
	c := b + 1
	for ; a <= '7'; a++ {
		for ; b <= '8'; b++ {
			for ; c <= '9'; c++ {
				if a != '7' {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
