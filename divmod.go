package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = 0
	*mod=0
*div=a / b
*mod=a % b
}
