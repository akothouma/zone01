package main

import ("github.com/01-edu/z01"
"os")

func printStr(s string){ 
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
func even(nbr int)bool{
	if nbr %2==0{
		return true
	}else{
		return false
	}
}

func isEven(nbr int) bool {
	if even(nbr) == true {
		return true
	} else {
		return false
	}
}

func main() {

	OddMsg:="I have an odd number of arguments"
	EvenMsg:="I have an even number of arguments"
	lengthOfArg:=len(os.Args[1:])
	if isEven(lengthOfArg) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
