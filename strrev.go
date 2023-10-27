package piscine

func StrRev(s string) string {
	var str2 string
	for i := range s {
		str2 += string(s[len(s)-i-1])
	}
	return str2
}
