package piscine

func LoafOfBread(str string) string {
	var requiredstring string
	for i := 0; i < len(str); i++ {
		requiredstring += str[i : i+4]
	}
	return requiredstring
}
