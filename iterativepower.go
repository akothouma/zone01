package piscine

func IterativePower(nb int, power int) int {
	var result int
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else {
		for i := power; i > 1; i-- {
			result = nb
			result = result * nb
		}
		return result
	}
}
