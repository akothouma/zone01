package piscine

func JumpOver(str string) string {
	if len(str) == 0 {
		return "\n"
	} else if len(str) < 3 {
		return "\n"
	}
	var result string
	for i := 0; i < len(str); i += 3 {
		result += string(str[i])
	}
	result += "\n"
	return result
}
