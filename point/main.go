package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func PrintString(s string) {
	for _, str := range s {
		z01.PrintRune(str)
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintNumRune(i int) {
	char := '0'
	for j := 0; j < i; j++ {
		char++
	}
	z01.PrintRune(char)
}

func main() {
	points := &point{}

	setPoint(points)
	a := points.x / 10
	b := points.x % 10
	c := points.y / 10
	d := points.y % 10

	PrintString("x = ")
	PrintNumRune(a)
	PrintNumRune(b)
	PrintString(", y = ")
	PrintNumRune(c)
	PrintNumRune(d)
	z01.PrintRune('\n')
}
