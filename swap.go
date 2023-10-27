package piscine

func Swap(a *int, b *int) {
	tempPointer := *a
	*a = *b
	*b = tempPointer
}
