package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	slicewanted := make([]int, 3)
	for i := 0; i < len(deck)/len(slicewanted); i++ {
		b := 0
		slicewanted = deck[b : b+len(slicewanted)]
		b += 3
		i++
		fmt.Printf("Player %v:%v", i+1, slicewanted)
	}
}
