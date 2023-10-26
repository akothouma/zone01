package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	}

	z01.PrintRune('F')
}

func main() {
	z01.IsNegative(1)
	z01.IsNegative(0)
	z01.IsNegative(-1)
}
