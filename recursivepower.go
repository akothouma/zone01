package piscine

func RecursivePower(nb int, power int) int {
	var result int
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else {
		result = nb
		if nb < 0 && power%2 == 1 {
			result = -result * RecursivePower(nb, power-1)
		} else {
			result = result * RecursivePower(nb, power-1)
		}
	}
	return result
}
