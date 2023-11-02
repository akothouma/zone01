package piscine

func ToLower(s string) string {
	var loweredSentence string
	var char rune
	for _, str := range s {
		if str >= 'A' && str <= 'Z' {
			char = str + 32
			loweredSentence += string(char)
		}
	}
	return loweredSentence
}
