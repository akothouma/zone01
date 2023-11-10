package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	for index := 0; index < len(a); index++ {
		if a[index+1] < a[index] {
			a[index+1], a[index] = a[index], a[index+1]
		}
	}
	return a[len(a)-1]
}
