package piscine

func ActiveBits(n int) int {
	var wantedstring []int
	count := 0
	for n > 0 {
		remainder := n % 2
		n = n / 2
		wantedstring = append(wantedstring, remainder)
	}
	for i := 0; i < len(wantedstring); i++ {
		if wantedstring[i] == 1 {
			count++
		}
	}
	return count
}
