package piscine

import (
	"fmt"
)

func ShoppingSummaryCounter(str string) map[string]int {
	words := ""
	var requiredwords []string
	wordcount := make(map[string]int)
	for _, ch := range str {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			if words != "" {
				requiredwords = append(requiredwords, words)
				words = ""
			}
		} else {
			words += string(ch)
		}
		for _, word := range requiredwords {
			_, exists := wordcount[word]
			if exists {
				wordcount[word] += 1
			} else {
				wordcount[word] = 1
			}
		}
	}

	for key, value := range wordcount {
		fmt.Printf(key, "=>", value)
	}
	return wordcount
}
