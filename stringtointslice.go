package piscine

func StringToIntSlice(str string) []int {
	var requiredslice []int
	if len(str) == 0 {
		return []int(nil)
	}
	for _, char := range str {
		requiredslice = append(requiredslice, int(char))
	}
	return requiredslice
}
