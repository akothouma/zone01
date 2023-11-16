package main

import "fmt"

var char1 rune
var answer, result string

func Rot14(s string) string {
	for _, char := range s {
		if char >= 97 && char <= 122 || char >= 65 && char <= 90 {
			if char < 110 || char < 78 {
				char1 = char + 14
			} else if char > 110 || char > 78 {
				char1 = char - 12
			} else if char == 110 || char == 78 {
				char1 = char + 12
			}
		} else {
			char1 = char
		}
		answer += string(char1)
	}
	return answer
}

func main() {
	str := "Hello! How are You?"
	result = Rot14(str)
	fmt.Println(result)
}
