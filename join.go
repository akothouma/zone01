package piscine

func Join(strs []string, sep string) string {
	separator := []rune(sep)
	var joinedstring string
	for _, str := range strs {
		for _, char := range str {
			if char == ',' {
				char = separator[0]
			}
			joinedstring += string(char)
		}
	}
	return joinedstring
}
