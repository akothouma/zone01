package main

import "github.com/01-edu/z01"


func DescendComb() {
   for ch:='9';ch>='0';ch--{
	for chr:='9',chr>='0'; ch--
		z01.PrintRune(ch)
		z01.PrintRune(chr)
		if ch!==0{
			z01.PrintRune(',')
			z01.PrintRune(' ')
		   }
	}
	
   }
   func main(){
	DescendComb()
   }
   
}
