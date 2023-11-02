package piscine

func IsPrintable(s string) bool {
	for _, letter := range s {
		if letter >= 32 && letter <= 127 {
			continue
		} else {
			return false
		}
	}
	return true
}
