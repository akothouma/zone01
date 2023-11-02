package piscine

func ToUpper(s string) string {
	var character rune
	var str string
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			character = ch - 32
		}
		str = str + string(character)
	}
	return str
}
