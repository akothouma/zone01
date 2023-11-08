package main

import (
	"fmt"
	"os"
)

func main() {
	arguements := os.Args

	if len(arguements) <= 1 {
		fmt.Println("File name missing")
	} else if len(arguements) == 2 {

		file, err := os.Open("quest8.txt")

		if err != nil {
			fmt.Println(err.Error())
		} else {
			stringToPrint := make([]byte, 14)
			_, err := file.Read(stringToPrint)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(string(stringToPrint))
			}
			file.Close()
		}
	} else {
		fmt.Println("Too many arguments")
	}
}
