package piscine

import (
	"github.com/01-edu/z01"
)

func Compact(ptr *[]string) int {
	count := 0
	for _, char := range *ptr {
		if !(char == "") {
			count++
			for _, ch := range char {
				z01.PrintRune(ch)
			}
			z01.PrintRune('\n')
		}
	}
	return count
}
