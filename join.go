package piscine

func Join(strs []string, sep string) string {
	var joinedstring string
	for _, str := range strs {
		for _, char := range str {
			if char == ',' {
				char = ':'
			}
			joinedstring += string(char)
		}
	}
	return joinedstring
}
