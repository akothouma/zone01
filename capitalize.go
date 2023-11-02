package piscine

func Capitalize(s string) string {
	s = ToLower(s)
	for i, r := range s {
		if i == 0 {
			s = ToUpper(string(r)) + s[i+1:]
		} else {
			if IsAlpha(string(r)) && !IsAlpha(string(s[i-1])) {
				if i != len(s) {
					s = s[:i] + ToUpper(string(r)) + s[i+1:]
				} else {
					s = s[:i] + ToUpper(string(r))
				}
			}
		}
	}
	return s
}
