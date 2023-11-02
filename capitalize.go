package piscine

import "unicode"

func Capitalize(s string) string {
	var output []rune // create an output slice
	isWord := true
	for _, val := range s {
		if isWord && unicode.IsLetter(val) { // check if character is a letter convert the first character to upper case
			output = append(output, unicode.ToUpper(val))
			isWord = false
		} else if !unicode.IsLetter(val) {
			isWord = true
			output = append(output, val)
		} else {
			output = append(output, val)
		}
	}
	return string(output)
}
