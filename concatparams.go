package piscine

func ConcatParams(args []string) string {
	var requiredstring string
	for index, onestring := range args {
		if index == 0 {
			requiredstring = args[0]
		} else {
			requiredstring += "\n" + onestring
		}
	}
	return requiredstring
}
