package piscine

func NRune(s string, n int) rune {
	castedstring := []rune(s)
	index := len(s)
	if n <= index && n < 0 {
		return castedstring[n-1]
	}
	return 0
}
