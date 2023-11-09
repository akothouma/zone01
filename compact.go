package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, char := range *ptr {
		if char != "" {
			count++
			(*ptr)[count] = char
		}
	}
	(*ptr) = (*ptr)[:count]
	return count
}
