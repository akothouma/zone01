package piscine

func Sqrt(nb int) int {
	if nb%2 == 0 {
		for i := 1; i*i <= nb; i++ {
			if i*i == nb {
				return i
			} else {
				return 0
			}
		}
	}
	return 0
}
