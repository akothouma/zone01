package main

import ("github.com/01-edu/z01"
"os")

func printStr(s string) {
	str:=[]rune(s)
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr %2== 0 {
		return true
	} else {
		return false
	}
}

func main() {
	Arg:=os.Args[1:]
	lengthOfArg:=len(Arg)
	if isEven(lengthOfArg) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
