
package piscine
import "github.com/01-edu/z01"

func PrintStr(s string) {
	var stringRune =[]rune(s)

length:=len(s)-1

for i:=0 ;i <= length; i++{
   z01.PrintRune(stringRune[i])
}
}
