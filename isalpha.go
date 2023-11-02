package piscine

func IsAlpha(s string) bool {
	stringlength := len(s)
	for _, character := range s {
		if stringlength == 0 || character >= 65 && character <= 90 || character >= 97 && character <= 122 {
			continue
		} else {
			return false
		}
	}
	return true
}
