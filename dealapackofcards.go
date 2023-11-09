package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	b := 0
	count := 0
	if count < 3 {
		slicewanted := deck[b : b+3]
		b += 3
		fmt.Printf("Player %v:%v", count, slicewanted)
		count++
	}
}
