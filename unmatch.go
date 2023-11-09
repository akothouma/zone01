package piscine

func Unmatch(a []int) int {
	i:=0
	b:=0
	returnvalue:=0
	for i< len(a)/2+1 {
		
		if len(a[b:b+2])%2==0{
			b=b+2
		}else{
			returnvalue=a[b]
		}
		i++
	}
	 return returnvalue
	}
 