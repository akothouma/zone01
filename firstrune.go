package piscine

func FirstRune(s string) rune {
	var str rune
	for i := 0; i <= len(s)-1; i++ {
		if i == 0 {
			str = rune(s[0])
		}
	}
	return str
}
