package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var castNumber string = string(n)

	z01.PrintRune('"')
	for i := 0; i <= len(castNumber)-1; i++ {
		runecharacter := rune(castNumber[i])
		z01.PrintRune(runecharacter)
	}
	z01.PrintRune('"')
}
