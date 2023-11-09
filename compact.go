package piscine

import "fmt"

func Compact(ptr *[]string) int {
	count := 0
	for _, char := range *ptr {
		if !(char == "") {
			count++
			fmt.Print(char)
			fmt.Print('\n')
		}
	}
	return count
}
