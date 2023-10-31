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
		result = nb
		for i := power; i > 1; i-- {
			if nb < 0 && power%2 == 1 {
				result = -(result * nb)
			} else {
				result = result * nb
			}
		}
		return result
	}
}
