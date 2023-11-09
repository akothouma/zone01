package piscine
import "github.com/01-edu/z01"


func Rot14(s string) string {
  var toPrint string
  for _,char:= range s{
	if char >='a' && char <='z' || char>='A' && char<='Z'{
		char=char-10
	}
	z01.PrintRune(char)
	toPrint+=string(char)
}
return toPrint
}


