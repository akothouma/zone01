package piscine

import (
	"fmt"
	"strings"
)

func ShoppingSummaryCounter(str string) map[string]int {
	words := strings.Split(str, " ")
	wordcount := make(map[string]int)

	for _, word := range words {
		_, exists := wordcount[word]
		if exists {
			wordcount[word] += 1
		} else {
			wordcount[word] = 1
		}
	}

	for key, value := range wordcount {
		fmt.Printf(key, "=>", value)
	}
	return wordcount
}
