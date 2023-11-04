package piscine

import "fmt"

func ConcatParams(args []string) string {
	var requiredstring string
	for _, onestring := range args {

		requiredstring = onestring
		fmt.Print('\n')
	}
	return requiredstring
}
