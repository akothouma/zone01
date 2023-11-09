package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, char := range *ptr {
		if char != "" {

			(*ptr)[count] = char
			count++
		}
	}
	(*ptr) = (*ptr)[:count]
	return count
}
