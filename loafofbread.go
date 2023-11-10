package piscine

func LoafOfBread(str string) string {
	lengthofslice := 5
	var requiredstring string
	for i := 0; i < len(str)-1; i += lengthofslice + 1 {
		requiredstring += str[i:i+lengthofslice] + " "
	}
	return requiredstring
}
