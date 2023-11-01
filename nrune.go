package piscine

func NRune(s string, n int) rune {
	castedstring := []rune(s)
	index := len(s)
	if n <= index {
		return castedstring[n-1]
	}
	return 0
}
