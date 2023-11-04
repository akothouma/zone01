package piscine

func ConcatParams(args []string) string {
	var requiredstring string
	for _, onestring := range args {
		requiredstring += onestring + "\n"
	}
	return requiredstring
}
