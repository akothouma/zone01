package piscine

func IsAlpha(s string) bool {
	stringlength := len(s)
	if stringlength == 0 {
		return true
	}
	for _, character := range s {
		if character >= 65 && character <= 90 || character >= 97 &&
			character <= 122 || character >= '0' && character <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
