package main
import "github.com/01-edu/z01"

func printStr(s string) {
	str:=[]rune(s)
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool{
	if nbr %2 == 1 {
		return false
	} else {
		return true
	}
}

func main() {
	if isEven(4){
		printStr("EvenMsg")
	} else {
		printStr("OddMsg")
	}
}
