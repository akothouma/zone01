package piscine

func LastRune(s string) rune {
	castedstring := []rune(s)
	return castedstring[len(castedstring)-1]
}
