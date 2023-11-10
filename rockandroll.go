package piscine

func RockAndRoll(n int) string {
	stringToDisplay := "error: number is negative\n"
	divisibleby2 := "rock\n"
	divisibleby3 := "roll\n"
	Nondiv := "error: non divisible\n"
	both := "rock and roll\n"

	if n < 0 {
		return stringToDisplay
	} else if n%2 == 0 {
		return divisibleby2
	} else if n%3 == 0 {
		return divisibleby3
	} else if n%2 == 0 && n%3 == 0 {
		return both
	} else {
		return Nondiv
	}
}
