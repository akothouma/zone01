package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for a := '0'; a <= '9'; a++ {
			for b := '0'; b <= '9'; b++ {
				for j := '0'; j <= '9'; j++ {
					if a < i || (a == i && b < j) {
						z01.PrintRune(i)
						z01.PrintRune(a)
						z01.PrintRune(' ')
						z01.PrintRune(j)
						z01.PrintRune(b)
						if a != '9' || b != '8' || i != '9' || j != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
