package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if s[i:i+len(toFind)-1] == toFind {
			return i
		}
	}
	return -1
}
