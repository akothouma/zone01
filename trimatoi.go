package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	intsrt := ""
	for _, item := range s {
		if item == '-' && len(intsrt) == 0 {
			intsrt += "-"
		}
		if item < '0' || item > '9' {
			continue
		}
		intsrt += string(item)
	}
	if len(intsrt) == 0 || intsrt == "-" {
		return 0
	}
	return Atoi(intsrt)
}
