package main

import (
	"fmt"
	"os"
)

func main() {
	arguements := os.Args[1:]
	for _, str := range arguements {
		if str == "01" || str == "galaxy" || str == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
