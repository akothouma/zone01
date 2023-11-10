package piscine

func LoafOfBread(str string) string {
	var requiredstring string
	for i := 0; i < len(str); i += 4 {
		requiredstring = append(requiredstring, str[0])
	}
	return requiredstring
}
