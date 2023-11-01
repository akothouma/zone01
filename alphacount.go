package piscine

func AlphaCount(s string) int {
	var count int
	for _, character := range s {
		if character >= 65 && character <= 90 || character >= 97 && character <= 122 {
			count++
		}
	}
	return count
}
