package piscine

func SplitWhiteSpaces(s string) []string {
	var requiredstring []string
	word := ""
	for _, ch := range s {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			if word != "" {
				requiredstring = append(requiredstring, word)
				word = ""
			} else {
				word += string(ch)
			}
		}
	}
	if word != " " {
		requiredstring = append(requiredstring, word)
	}
	return requiredstring
}
