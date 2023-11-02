package piscine

func IsUpper(s string) bool {
	for _, character := range s {
		if character < 65 || character > 90 {
			return false
		}
	}
	return true
}
