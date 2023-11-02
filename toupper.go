package piscine

func ToUpper(s string) string {
	var character rune
	var str string
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			character = ch - 32
			str += string(character)
		} else {
			str += string(ch)
		}
	}
	return str
}
