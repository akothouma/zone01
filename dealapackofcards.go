package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	count := 1
	for b := 0; b < 12; b += 3 {

		fmt.Printf("Player %v: %v, %v, %v\n", count,
			deck[b], deck[b+1], deck[b+2])
		count++
	}
}
